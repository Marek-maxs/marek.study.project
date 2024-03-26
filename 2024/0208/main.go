package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/en-trak/mqclient/v3/amqp"
	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/proto"


	pb "github.com/en-trak/protobuf/v2/sbos/hierarchy"
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

func main () {
	// branch
	err := PublishPovUpdateSummaryEvent("f4e23142-bc87-4518-af93-739dc1396723")
	if err != nil {
		fmt.Println(err)
	}
	//// branch
	//err = PublishPovUpdateSummaryEvent("7ec46c86-85a0-4742-b806-3a5f5a81ddfc")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//// master
	//err = PublishPovUpdateSummaryEvent("7dfb6001-b687-4c44-ac75-1c49d149eb8a")
	//if err != nil {
	//	fmt.Println(err)
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

func PublishPovUpdateSummaryEvent(pov string) error {
	var startDate, endDate int64
	ctx := context.Background()
	povID := uuid.FromStringOrNil(pov)
	startDate, endDate = 1708147881, 1708617600

	msg := &pb.MQChangedPovUpdateSummary{
		PovId:     povID.Bytes(),
		StartDate: startDate,
		EndDate:   endDate,
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

	err = client.Publish(
		amqp.WithPublishContext(ctx),
		amqp.WithPublishPayload(payload),
		amqp.WithPublishTopic("event.hierarchy.pov_changed"),
	)

	if err != nil {
		return fmt.Errorf("publish failed on topic %s, %s",
			"event.hierarchy.pov_changed", err.Error())
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
		"192.168.83.128",
		"guest",
		"guest",
		5672,
		false,
	)
}