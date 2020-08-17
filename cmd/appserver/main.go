package main

import (
	. "client-go-demo/config"
	"client-go-demo/pkg/middleware"
	. "client-go-demo/pkg/util"
	"os"
)

func main() {
	e := middleware.SetupGin()

	conf := GetConfigFromENV()

	Logger.SetLogLevel(conf.LogLevel)

	client, err := conf.Kube.New()
	if err != nil {
		Logger.Errorf("main", "get cluster client error:%s", err.Error())
		os.Exit(1)
	}

	Register(e, client)

	addr := conf.Server.GetServerAddr()
	err = e.Run(addr)
	if err != nil {
		Logger.Errorf("main", "run server error:%s", err.Error())
		os.Exit(1)
	}
}
