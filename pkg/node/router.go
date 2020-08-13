package node

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func NodeRouter(client *kubernetes.Clientset, c *gin.RouterGroup) {
	nRouter := c.Group("/nodes")
	{
		nRouter.GET("", func(c *gin.Context) {
			ListNodeHandler(client, c)
		})

		nRouter.GET("/:name", func(c *gin.Context) {
			GetNodeHandler(client, c)
		})

		nRouter.GET("/:name/pods", func(c *gin.Context) {
			ListPodsOfNodeHandler(client, c)
		})

		nRouter.GET("/:name/events", func(c *gin.Context) {
			ListEventsOfNodeHandler(client, c)
		})

		nRouter.PUT("/:name", func(c *gin.Context) {
			UpdateNodeHandler(client, c)
		})
	}
}
