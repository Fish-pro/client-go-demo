package config

import (
	"client-go-demo/pkg/app/deployment"
	"client-go-demo/pkg/app/node"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func Register(r *gin.Engine, client *kubernetes.Clientset) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	v1Router := r.Group("/v1")

	// deployment router group
	deployment.DeployRouter(client, v1Router)

	// node router group
	node.NodeRouter(client, v1Router)

}
