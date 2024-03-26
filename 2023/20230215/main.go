package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

/**
*
* Author: Marek
* Date: 2023-02-15 15:35
* Email: 364021318@qq.com
*
 */

func main() {
	go func() {
		for {
			log.Info().Msg("check log")
			now := time.Now()
			next := now.Add(time.Minute * 2)
			next = time.Date(next.Year(), next.Month(), next.Day(), 16,0,30,30, next.Location())
			t := time.NewTicker(next.Sub(now))
			<-t.C
		}
	}()
	time.Sleep(time.Minute * 10)
	//t := time.NewTimer(time.Second * 10)
	//for {
	//	select {
	//	case <-t.C:
	//		t.Reset(time.Second* 10)
	//		fmt.Println("print reset time")
	//	}
	//}
}