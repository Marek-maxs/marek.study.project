package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/en-trak/mqclient/v3"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	"math"
	"time"
)

type bmsmeterData struct {
	Timestamp     uint32  `json:"Timestamp"`
	DatapointName string  `json:"DatapointName"`
	M3           float32 `json:"M3"`
	Serial        string  `json:"Serial"`
}

func main() {
clien, err := NewPubMqtt("hub/bc33674b-0cac-41a3-9d69-16d121619343/bmsWaterMeterData")
if err != nil{
	log.Fatal().Err(err).Msg("connect mqtt failed")
}

data := bmsmeterData{
	Timestamp:     uint32(time.Now().Unix()),
	DatapointName: "test-datapoint-name",
	M3:           100.9,
	Serial:        "b8:27:eb:6e:d4:75",
}
	payload, err := json.Marshal(data)
if err != nil {
	log.Fatal().Err(err).Msg("marshal json failed")
}

opt := []mqclient.PublishOption{mqclient.WithPublishPayload(payload)}
opt = append(opt, mqclient.WithPublishQos(mqclient.QosOne))

err = clien.Publish(opt...)
if err != nil {
	log.Fatal().Err(err).Msg("publish mqtt")
}
}

func test() {
	data := "9125435e"
	data1, _ := hex.DecodeString(data)
	d := math.Float32frombits(binary.BigEndian.Uint32(data1))
	fmt.Println(d)
	loc, _ := time.LoadLocation("Asia/Hong_Kong")
	localTime := time.Now().In(loc)
	endTime := time.Date(localTime.Year(), localTime.Month(), localTime.Day(),
		0, 0, 0, 0, loc)

	fmt.Println(endTime.Unix())

	a := 14
	b := 30
	c := 14.28
	d = float32(a) + float32(b)/100
	if float32(c) > d {
		fmt.Println("333232342342")
	}

	fmt.Println(float32(b)/100)

	old := "2022-05-30 14:30:00"
	new := "2022-05-30 13:28:00"
	oldTime, _ := time.Parse("2006-01-02 15:04:05", old)
	newTime, _ := time.Parse("2006-01-02 15:04:05", new)
	fmt.Println("oldTime时间是否在newTime时间之后：", oldTime.After(newTime))
	fmt.Println("oldTime时间是否在newTime时间之前：", oldTime.Before(newTime))
	//fmt.Println("oldTime时间是否等于newTime时间：", oldTime.Equal(newTime))
	nowDate := time.Now().Format("2006-01-02")
	fmt.Println(fmt.Sprintf("%s %d:%d:00", nowDate, 8, 00))
	Delta1, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s:%s:00", nowDate, "8", "00"))
	fmt.Println(Delta1)

	//recore := "2023-10-10 15:35:40"
	//recoreTime, _ := time.Parse("2006-01-02 15:04:05", recore)
	//// delta2, _ := time.Parse("2006-01-02 15:04:05", "2023-10-11 09:30:00")
	//delta4, _ := time.Parse("2006-01-02 15:04:05", "2023-10-10 14:30:00")

	recoredTime := 1696919883

	delta1 := 1696924800
	delta2 := 1696930200
	delta3 := 1696941000
	delta4 := 1696948200

	fmt.Println(delta4 <= recoredTime)
	fmt.Println(time.Unix(int64(recoredTime), 0))
	fmt.Println(time.Unix(int64(delta4), 0))

	//loc,_ := time.LoadLocation("Asia/Hong_Kong")
	delta4Time, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s:%s:00", "2023-10-10",
		fmt.Sprintf("%d", 14), fmt.Sprintf("%d", 30)), loc)

	delta4 = int(delta4Time.Unix())

	//fmt.Println(time.Unix(delta44, 0))

	switch {
	case delta4 <= recoredTime:
		fmt.Println("delta4")
	case delta3 <= recoredTime && delta4 > recoredTime:
		fmt.Println("delta3")
	case delta2 <= recoredTime && delta3 > recoredTime:
		fmt.Println("delta2")
	case delta1 <= recoredTime && delta2 > recoredTime:
		fmt.Println("delta1")
	default:
		fmt.Println("not match")
		return
	}

	meterID := uuid.Must(uuid.NewV4())

	meterIDStr := base64.StdEncoding.EncodeToString(meterID.Bytes())
	fmt.Println(meterIDStr)
	isAirconOperateing := false

	if !isAirconOperateing {
		fmt.Println("aircon operation")
	}
}

func NewPubMqtt(topic string) (mqclient.Publisher, error) {
	return mqclient.NewPublisher(
		mqclient.WithBlock(),
		mqclient.WithDefaultTopic(topic),
		mqclient.WithBroker(InitMqtt()),
	)
}

func InitMqtt() mqclient.Broker {
	var tls *mqclient.TLS

	return mqclient.NewMQTTBroker(
		"dev.mqtt.en-trak.com",
		"hub",
		"mvhsOzg8NTwxXYQkJk",
		1883,
		tls,
	)
}