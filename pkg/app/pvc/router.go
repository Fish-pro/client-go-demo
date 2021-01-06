package pvc

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func PersistentVolumeClaimRouter(client *kubernetes.Clientset, c *gin.RouterGroup) {
	pRouter := c.Group("/namespace/:namespace/pvc")
	{
		pRouter.GET("", func(c *gin.Context) {
			listPersistentVolumeClaim(client, c)
		})
		pRouter.GET("/:name", func(c *gin.Context) {
			getPersistentVolumeClaim(client, c)
		})
		pRouter.POST("", func(c *gin.Context) {
			createPersistentVolumeClaim(client, c)
		})
		pRouter.PUT("/:name", func(c *gin.Context) {
			updatePersistentVolumeClaim(client, c)
		})
		pRouter.DELETE("/:name", func(c *gin.Context) {
			deletePersistentVolumeClaim(client, c)
		})
	}
}
