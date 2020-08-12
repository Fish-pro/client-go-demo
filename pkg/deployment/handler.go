package deployment

import (
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListDeployHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")

	deploys, err := client.AppsV1().Deployments(namespace).List(metav1.ListOptions{})
	if err != nil {
		c.JSON(400, "get deploys error")
		return
	}
	c.JSON(200, deploys)
}

func GetDeployHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")
	name := c.Param("name")

	deploy, err := client.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		c.JSON(400, "get deploy error")
		return
	}

	c.JSON(200, deploy)
}
