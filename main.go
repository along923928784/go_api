package main

import (
	"errors"
	"time"

	"net/http"

	"github.com/lexkong/log"
	"jiyue.im/conf"
	"jiyue.im/server"
)

func main() {
	var addr string

	conf.Init(&addr)

	r := server.LoadRouter()

	// 生死检测
	go func() {
		if err := pingServer(addr); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", addr)
	log.Info(http.ListenAndServe(addr, r).Error())
}

// 健康检测
func pingServer(addr string) error {
	for i := 0; i < 5; i++ {
		resp, err := http.Get("http://" + addr + "/ping")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
