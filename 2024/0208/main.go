package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strconv"

	"github.com/en-trak/mqclient/v3/amqp"
	"github.com/gofrs/uuid"
)

/**
*
* Author: Marek
* Date: 2024-02-08 10:47
* Email: 364021318@qq.com
*
* 模拟 hierarchy 向rabbitmq 发送 pov 的更新消息
*
 */

var commonMQPublisher *SmartPublisher

var errCommonMQPublisherEmpty = errors.New("empty common mq publisher")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Invalid arguments")
		return
	}
	id := args[0]

	startTime := args[1]
	endTime := args[2]

	if startTime == "" || endTime == "" {
		fmt.Println("Invalid arguments startTime endTime")
		return
	}

	st, _ := strconv.ParseInt(startTime, 0, 64)
	et, _ := strconv.ParseInt(endTime, 0, 64)

	// branch
	err := PublishPovUpdateSummaryEvent(id, st, et)
	if err != nil {
		fmt.Println(err)
	}
}

func initCommonMQPublisher() error {
	if commonMQPublisher != nil {
		return nil // already init, do nothing here
	}

	cli, err := NewSmartPublisher()
	if err != nil {
		return err
	}

	commonMQPublisher = cli
	return nil
}

func getCommonMQPublisher() (amqp.Publisher, error) {
	if commonMQPublisher == nil {
		initCommonMQPublisher()
	}

	return commonMQPublisher.Cli()
}

func PublishPovUpdateSummaryEvent(pov string, startTime, endTime int64) error {
	ctx := context.Background()
	povID := uuid.FromStringOrNil(pov)

	//msg := &pb.MQChangedPovUpdateSummary{
	//	PovId:     povID.Bytes(),
	//	StartDate: startDate,
	//	EndDate:   endDate,
	//	Type:      pb.PovUnits_P_KWH,
	//}

	client, errClient := getCommonMQPublisher()
	if errClient != nil {
		return errClient
	}

	//payload, err := proto.Marshal(msg)
	//if err != nil {
	//	return err
	//}
	payload := []byte(
		fmt.Sprintf(`{"PovID":"%s", "StartTime": %d, "EndTime": %d,"Timezone": "%s", "Data":""}`,
			povID, startTime, endTime, "Asia/Hong_Kong"))
	err := client.Publish(
		amqp.WithPublishContext(ctx),
		amqp.WithPublishPayload(payload),
		amqp.WithPublishTopic("energySummaryTopic"),
	)

	if err != nil {
		return fmt.Errorf("publish failed on topic %s, %s",
			"energySummaryTopic", err.Error())
	}

	return nil
}

var errSmartPublisherClosed = errors.New("MQ publisher is closed")

type SmartPublisher struct {
	cli      amqp.Publisher
	isClosed bool
}

func NewSmartPublisher() (*SmartPublisher, error) {
	cli, err := NewPublisher()
	if err != nil {
		return nil, err
	}

	return &SmartPublisher{
		cli: cli,
	}, nil
}

func (s *SmartPublisher) Cli() (amqp.Publisher, error) {
	if s.isClosed {
		return nil, errSmartPublisherClosed
	}

	return s.cli, nil
}

func (s *SmartPublisher) Close() error {
	if s.isClosed {
		return nil
	}

	s.isClosed = true
	return s.cli.Close()
}

func NewPublisher() (amqp.Publisher, error) {
	return amqp.NewPublisher(
		amqp.WithBlock(),
		amqp.WithBroker(ParseAMQPBroker()),
	)
}

func ParseAMQPBroker() amqp.Broker {
	return amqp.NewBroker(
		"127.0.0.1",
		"guest",
		"guest",
		5672,
		false,
	)
}
