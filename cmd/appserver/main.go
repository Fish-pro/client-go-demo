package main

import (
	. "client-go-demo/config"
	"client-go-demo/pkg/middleware"
	. "client-go-demo/pkg/util"
	"os"
)

func main() {
	router := middleware.SetupGin()

	conf := GetConfigFromENV()

	Logger.SetLogLevel(conf.LogLevel)

	client, err := conf.GetClusterClient()
	if err != nil {
		Logger.Errorf("main", "get cluster client error:%s", err.Error())
		os.Exit(1)
	}

	Register(router, client)

	err = router.Run(conf.GetServerAddr())
	if err != nil {
		Logger.Errorf("main", "run server error:%s", err.Error())
		os.Exit(1)
	}
}
