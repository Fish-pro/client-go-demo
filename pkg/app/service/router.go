package service

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func SvcRouter(client *kubernetes.Clientset, c *gin.RouterGroup) {
	sRouter := c.Group("/services")
	{
		sRouter.GET("", func(c *gin.Context) {
			ListServiceHandler(client, c)
		})

		sRouter.GET("/:name", func(c *gin.Context) {
			GetServiceHandler(client, c)
		})

		sRouter.POST("", func(c *gin.Context) {
			CreateServiceHandler(client, c)
		})

		sRouter.PUT("/:name", func(c *gin.Context) {
			UpdateServiceHandler(client, c)
		})

		sRouter.DELETE("/:name", func(c *gin.Context) {
			DeleteServiceHandler(client, c)
		})

		sRouter.GET("/:name/events", func(c *gin.Context) {
			ListEventsOfServiceHandler(client, c)
		})

		sRouter.GET("/:name/pods", func(c *gin.Context) {
			ListPodsOfServiceHandler(client, c)
		})
	}
}
