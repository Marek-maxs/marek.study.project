package main

import (
	"fmt"
	"sync"
)

type Event struct {
	eventMqMeterData        MqMeterData
}

type MqMeterData struct{
	*grpcEnergy
}

func (m *MqMeterData) ParserMq() {
	fmt.Println(m.limit)
}

type grpcEnergy struct {
	mutex sync.Mutex
	//pool  chan *grpc.ClientConn
	host  string
	total int
	limit int
}

func NewGrpcEnergy() *grpcEnergy {
	grpc := &grpcEnergy{
		mutex: sync.Mutex{},
		host:  "www.baidu.com",
		total: 100,
		limit: 220,
	}
	return grpc
}

func main()  {
	newEvent := Event{}
	newEvent.eventMqMeterData.grpcEnergy = NewGrpcEnergy()
	newEvent.eventMqMeterData.ParserMq()
}