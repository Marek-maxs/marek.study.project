package rabbitmq

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/streadway/amqp"
)

// 测试用例
type Obj struct {
	Item1 string `json:"Data"`
	Item2 string `json:"item2"`
	Item3 string `json:"item3"`
}

func StartAMQPConsume() {
	defer func() {
		if err := recover(); err != nil {
			time.Sleep(3 * time.Second)
			fmt.Println("休息3秒")
			StartAMQPConsume()
		}
	}()
	conn, err := amqp.Dial("amqp://guest:guest@192.168.83.128:5672/")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()
	closeChan := make(chan *amqp.Error, 1)
	notifyClose := ch.NotifyClose(closeChan)
	closeFlag := false
	msgs, err := ch.Consume(
		"sub_ems_energy_bmsMeterPayloadTopic",
		"",
		false,
		false,
		false,
		false,
		nil,
		)
	var obj Obj
	for {
		select {
		case e := <-notifyClose:
			fmt.Println("chan通道错误，e: %s", e.Error())
			close(closeChan)
			time.Sleep(5 * time.Second)
			StartAMQPConsume()
			closeFlag = true
		case msg := <-msgs:
			if err := json.Unmarshal(msg.Body, &obj); err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(obj.Item1)
			//msg.Ack(true)
		}
		if closeFlag {
			break
		}
	}
}