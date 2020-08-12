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
	}
}
