package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"errors"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	dac "github.com/xinsnake/go-http-digest-auth-client"
)

/**
*
* Author: Marek
* Date: 2023-05-16 12:32
* Email: 364021318@qq.com
*
 */

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

func main() {
	data, err := GetXMLData("en-trak1215.d.en-trak.com", time.Now().Add(-160*time.Minute).Unix(), 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	m := XMLSchema{}

	err = xml.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
	}
	val := m.DataAttr.Timestamp[2:]
	n, err := strconv.ParseUint(val, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	ts := int64(n)
	datapointList := m.DataAttr.RegMap()
	if len(m.DataAttr.Records) == 0 {
		fmt.Println("meter cannot be obtained at the time:")
	}
	for _, oneRecord := range m.DataAttr.Records{
		for datapointName, regItem := range datapointList{
			if regItem.Type != "P" {
				continue
			}
			trimDatapointName := strings.TrimSpace(datapointName)
			if trimDatapointName == "DB-FP" {
				fmt.Println(trimDatapointName)
				fmt.Println(float32(oneRecord.Values[regItem.Idx]/3600))
				fmt.Println(time.Unix(ts, 0).Format("2006-01-02 15:04"))
			}

		}
	}
	fmt.Println(m)
}

const BMSUSERNAME   = "entrak"
// getXMLData get data of meter by url
func GetXMLData(url string, timestamp int64, beforeMins int) ([]byte, error) {
	if url == "" {
		return nil, errors.New("xml_url is empty")
	}

	username := BMSUSERNAME
	password := "jianshu1906"
	t := dac.NewTransport(username, password)
	url = fmt.Sprintf(`http://%s/cgi-bin/egauge-show/?m&n=%d&a&f=%d`, url, beforeMins, timestamp)
	// avoid 'dial tcp: too many open files' error, add timeout for each request
	http.DefaultClient.Timeout = time.Second * 180
	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Debug().Err(err).Msg(fmt.Sprintf("[ContextDeadlineExceeded] egauge: %s\n", url))
		}

		if os.IsTimeout(err) {
			log.Debug().Err(err).Msg(fmt.Sprintf("[TimeoutError] egauge: %s\n", url))
		}
		// log.Error().Msg(fmt.Sprintf("%s", err))
		return nil, err
	}

	req.Header.Set("Accept", "text/xml; charset=utf-8")
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")

	resp, err := t.RoundTrip(req)
	if err != nil {
		log.Debug().Msg(fmt.Sprintf("%s", err))
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Debug().Msg(fmt.Sprintf("get_xml_err:%s", err))
			return nil, err
		}

		return bodyBytes, nil
	}

	log.Debug().Str("URL", url).Msg(`[failed] get egaugemeter data from xml`)

	return nil, fmt.Errorf("failed get egaugemeter data from xml")
}

func (d *DataAttr) RegMap() map[string]*IdxRegister {
	res := map[string]*IdxRegister{}

	for idx, r := range d.Registers {
		res[r.Name] = &IdxRegister{
			Idx:      idx,
			Register: r,
		}
	}

	return res
}