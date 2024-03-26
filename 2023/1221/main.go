package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

/**
*
* Author: Marek
* Date: 2023-12-21 10:45
* Email: 364021318@qq.com
* 获取https 证书信息与过期信息
 */

func main() {
	address := "dev.mq2tt.en-trak.com:8883"
	getTLSInfoWithTCP(address)
}


func getTLSInfoWithTCP(address string) {
	conn, err := tls.Dial("tcp", address, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("过期时间：",conn.ConnectionState().PeerCertificates[0].NotAfter)
	fmt.Println("组织信息: ", conn.ConnectionState().PeerCertificates[0].Subject.CommonName)
	fmt.Println("success end")
}

func getHTTPSTLS() {
	tr := &http.Transport{TLSClientConfig:&tls.Config{InsecureSkipVerify:true}}

	client := &http.Client{Transport:tr}
	seedURL := "https://mqtt.en-trak.com"
	resp, err := client.Get(seedURL)
	if err != nil {
		fmt.Errorf(seedURL, "请求失败")
		panic(err)
	}

	defer resp.Body.Close()

	certInfo := resp.TLS.PeerCertificates[0]
	fmt.Println("过期时间：", certInfo.NotAfter)
	fmt.Println("组织信息: ", certInfo.Subject)
}