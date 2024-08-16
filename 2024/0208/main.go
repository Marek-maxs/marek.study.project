package main

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/en-trak/mqclient/v3/amqp"
	pb "github.com/en-trak/protobuf/v2/sbos/hierarchy"
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
	//flag.Parse()
	//args := flag.Args()
	//if len(args) == 0 {
	//	fmt.Println("Invalid arguments")
	//	return
	//}
	//id := args[0]
	//
	//startTime := args[1]
	//endTime := args[2]
	//
	//if startTime == "" || endTime == "" {
	//	fmt.Println("Invalid arguments startTime endTime")
	//	return
	//}
	//
	//st, _ := strconv.ParseInt(startTime, 0, 64)
	//et, _ := strconv.ParseInt(endTime, 0, 64)
	//st := int64(1726099200)
	//et := int64(1726102800)
	//
	//err := PublishPovUpdateSummaryEvent("bdd42b23-67a1-4ab4-a400-6dde3a3430e1", st, et)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = PublishPovUpdateSummaryEvent("815b8c12-853c-43ac-b864-bfd25d6b026f", st, et)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = PublishPovUpdateSummaryEvent("e5b4345e-20af-470f-b7d3-2cb774f5b788", st, et)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = PublishPovUpdateSummaryEvent("e5ad9e76-0f8b-41e4-9fdd-bbe57320cf7b", st, et)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = PublishPovUpdateSummaryEvent("7ba49fa2-2361-4d6d-93fb-400530973185", st, et)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = PublishPovUpdateSummaryEvent("dc8378f0-d870-47d5-b7df-f806793e9846", st, et)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = PublishPovUpdateSummaryEvent("4c02a037-b930-4619-86db-265268374f8a", st, et)
	//if err != nil {
	//	fmt.Println(err)
	//}

	st := int64(1730390400)
	et := int64(1730476799)
	err := PublishPovUpdateSummaryEvent("5714278a-db6a-4d74-99ec-c80829216429", st, et)
	if err != nil {
		fmt.Println(err)
	}
	//
	//old := st
	//// branch
	//i := 0
	//for {
	//	i++
	//	fmt.Println(i)
	//	st = et - 3600
	//	err := PublishPovUpdateSummaryEvent("b75b1e59-9a16-4ac5-a7d3-f96c892bcd18", st, et)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	et = st
	//	if old == st {
	//		break
	//	}
	//}
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

	msg := &pb.MQChangedPovUpdateSummary{
		PovId:     povID.Bytes(),
		StartDate: startTime,
		EndDate:   endTime,
		Type:      pb.PovUnits_P_LITRE,
	}

	client, errClient := getCommonMQPublisher()
	if errClient != nil {
		return errClient
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	payload = []byte(
		fmt.Sprintf(`{"PovID":"%s", "StartTime": %d, "EndTime": %d,"Timezone": "%s", "Data":""}`,
			povID, startTime, endTime, "Asia/Hong_Kong"))
	err = client.Publish(
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
