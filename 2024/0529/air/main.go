package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

/*

创建一个服务，可以通过8888端口进行访问的
传入index 参数可以得到 Hello world 的结果
主要使用 http.service 的包的自带的功能

*/

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func main() {
	log.Info().Msg("start run service")

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:                         ":8888",
		Handler:                      mux,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}

	log.Info().Msg("starting run service")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("listening on 8888 failed")
	}
}
