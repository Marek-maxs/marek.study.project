package main

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

// 开微信服务：

const (
	WxToken = "gyjJA4JdIwRiyY2S3i05RdNPKZb"
)

func main() {
	http.HandleFunc("/follow", weixinSer)
	http.ListenAndServe("127.0.0.1:8889", nil)
}

func weixinSer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !validateWechatRequest(w, r) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}
	if r.Method == "POST" {
		textRequestBody := parseTextRequestBody(r)
		if textRequestBody != nil {
		}
		fmt.Printf("Wechat Service: Recv text msg [%s] form user [%s]\n",
			textRequestBody.Content,
			textRequestBody.FromUserName,
		)
		responseTextBody, err := makeTextResponseBody(
			textRequestBody.ToUserName,
			textRequestBody.FromUserName,
			"Hello, "+textRequestBody.FromUserName,
		)
		if err != nil {
			log.Println("Wechat Service: makeTextResponseBody error: ", err)
			return
		}
		fmt.Fprint(w, string(responseTextBody))
	}
}

// 验证消息是否来自微信：

func validateWechatRequest(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()

	signature := r.FormValue("signature")

	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")

	echostr := r.FormValue("echostr")

	hashcode := makeSignature(WxToken, timestamp, nonce)

	log.Printf("Try validateWechatRequest: hashcode: %s, signature: %s\n", hashcode, signature)
	if hashcode == signature {
		fmt.Fprintf(w, "%s", echostr)
		return true
	} else {
		fmt.Fprintf(w, "hashcode != signature")
	}
	return false
}

func makeSignature(token, timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)

	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))

	return fmt.Sprintf("%x", s.Sum(nil))
}

// 微信消息解析：

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

func parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(string(body))
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

// 微信消息响应：

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

type CDATAText struct {
	Text string `xml:",innerxml"`
}

func value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(fromUserName)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(content)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}