package pod

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func PodRouter(client *kubernetes.Clientset, c *gin.RouterGroup) {
	pRouter := c.Group("/pods")
	{
		pRouter.GET("", func(c *gin.Context) {
			ListPodHandler(client, c)
		})
		pRouter.GET("/:name", func(c *gin.Context) {
			GetPodHandler(client, c)
		})
		pRouter.POST("", func(c *gin.Context) {
			CreatePodHandler(client, c)
		})
		pRouter.PUT("/:name", func(c *gin.Context) {
			UpdatePodHandler(client, c)
		})
		pRouter.DELETE("/:name", func(c *gin.Context) {
			DeletePodHandler(client, c)
		})
		pRouter.GET("/:name/events", func(c *gin.Context) {
			ListEventsOfPodHandler(client, c)
		})
	}
}
