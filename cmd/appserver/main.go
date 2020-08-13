package main

import (
	"client-go-demo/config"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	router := gin.Default()

	conf, err := clientcmd.BuildConfigFromFlags(
		"https://10.6.124.52:16443",
		"/Users/york/go/src/github.com/Fish-pro/client-go-demo/config/52.yaml",
	)

	client, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic("that maybe have problem")
	}

	config.Register(router, client)

	router.Run("0.0.0.0:9090")
}
