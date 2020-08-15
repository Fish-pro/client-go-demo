package main

import (
	. "client-go-demo/config"
	"client-go-demo/pkg/middleware"
	. "client-go-demo/pkg/util"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func main() {
	router := middleware.SetupGin()

	conf := GetConfigFromENV()

	Logger.SetLogLevel(conf.LogLevel)

	clusterConfig, err := clientcmd.BuildConfigFromFlags(
		conf.MasterUrl,
		conf.KubernetesConfigPath,
	)

	if err != nil {
		Logger.Errorf("main", "build cluster config error")
		os.Exit(1)
	}

	client, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		Logger.Errorf("main", "that maybe have problem")
		os.Exit(1)
	}

	Register(router, client)

	err = router.Run(conf.GetServerAddr())
	if err != nil {
		Logger.Errorf("main", "run server error:%s", err.Error())
		os.Exit(1)
	}
}
