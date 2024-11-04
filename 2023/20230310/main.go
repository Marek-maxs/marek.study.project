package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
	dac "github.com/xinsnake/go-http-digest-auth-client"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	BMSUSERNAME = "entrak"
	PASSWORD = "jianshu1906"
)

type XMLSchema struct {
	XMLName  xml.Name `xml:"group"`
	Serial   string   `xml:"serial,attr"`
	DataAttr DataAttr `xml:"data"`
}

type DataAttr struct {
	XMLName   xml.Name   `xml:"data"`
	Columns   int        `xml:"columns,attr"`
	Timestamp string     `xml:"time_stamp,attr"`
	Timedelta int        `xml:"time_delta,attr"`
	Epoch     string     `xml:"epoch,attr"`
	Registers []Register `xml:"cname"`
	Records   []Record   `xml:"r"`
}

type Register struct {
	XMLName xml.Name `xml:"cname"`
	Type    string   `xml:"t,attr"`
	Did     string   `xml:"did,attr"`
	Name    string   `xml:",chardata"`
}

type IdxRegister struct {
	Idx int
	Register
}

type Record struct {
	XMLName xml.Name  `xml:"r"`
	Values  []float64 `xml:"c"`
}

func GetXMLData(url string, timestamp int, beforeMins int) {
	if url == "" {
		log.Fatal().Msg("url is empty")
	}
	t := dac.NewTransport(BMSUSERNAME, PASSWORD)
	url = fmt.Sprintf(`http://%s/cgi-bin/egauge-show/?m&n=%d&a&f=%d`, url, beforeMins, timestamp)
	http.DefaultClient.Timeout = time.Second * 100
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("get egauge meter failed")
	}
	req.Header.Set("Accept", "text/xml; charset=utf-8")
	req.Header.Set("Connection", "close")
	resp, err := t.RoundTrip(req)
	if err != nil {
		log.Fatal().Err(err).Msg("roundtrip failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal().Err(err).Msg("ReadAll failed")
		}
		m := XMLSchema{}
		err = xml.Unmarshal(bodyBytes, &m)
		if err != nil {
			log.Fatal().Err(err).Msg("Unmarshal failed")
		}
		fmt.Println(m)
	}
}

func main() {
	//url := "en-trak1096.d.en-trak.com"
	//minutes := 10
	//endTimestamp := 1678428000
	//GetXMLData(url, endTimestamp, minutes)
}