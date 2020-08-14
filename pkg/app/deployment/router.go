package deployment

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func DeployRouter(client *kubernetes.Clientset, c *gin.RouterGroup) {
	dRouter := c.Group("/deployments")
	{
		dRouter.GET("", func(c *gin.Context) {
			ListDeployHandler(client, c)
		})
		dRouter.GET("/:name", func(c *gin.Context) {
			GetDeployHandler(client, c)
		})
		dRouter.POST("", func(c *gin.Context) {
			CreateDeployHandler(client, c)
		})
		dRouter.PUT("/:name", func(c *gin.Context) {
			UpdateDeployHandler(client, c)
		})
		dRouter.DELETE("/:name", func(c *gin.Context) {
			DeleteDeployHandler(client, c)
		})
	}
}
