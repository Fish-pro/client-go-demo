package main

import (
	"client-go-demo/pkg/deployment"
	"client-go-demo/pkg/node"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	v1Router := router.Group("/v1")

	config, err := clientcmd.BuildConfigFromFlags(
		"https://10.6.124.52:16443",
		"/Users/york/go/src/github.com/Fish-pro/client-go-demo/config/52.yaml",
	)

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("that maybe have problem")
	}

	// deployment router group
	deployment.DeployRouter(client, v1Router)

	// node router group
	node.NodeRouter(client, v1Router)

	router.Run("0.0.0.0:9090")
}
