package deployment

import (
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListDeployHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")

	deploys, err := client.AppsV1().Deployments(namespace).List(metav1.ListOptions{})
	if err != nil {
		c.JSON(500, "get deploys error")
		return
	}

	deploys.Kind = "List"
	deploys.APIVersion = "v1"
	for index := range deploys.Items {
		deploys.Items[index].APIVersion = "apps/v1"
		deploys.Items[index].Kind = "Deployment"
	}

	c.JSON(200, deploys)
}

func GetDeployHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")
	name := c.Param("name")

	deploy, err := client.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		c.JSON(500, "get deploy error")
		return
	}

	deploy.APIVersion = "apps/v1"
	deploy.Kind = "Deployment"

	c.JSON(200, deploy)
}

func CreateDeployHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")

	var deployInfo v1.Deployment
	if err := c.BindJSON(deployInfo); err != nil {
		c.JSON(400, "body error")
		return
	}

	deploy, err := client.AppsV1().Deployments(namespace).Create(&deployInfo)
	if err != nil {
		c.JSON(500, "create deploy error")
		return
	}

	c.JSON(200, deploy)
}

func UpdateDeployHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")
	name := c.Param("name")

	if name == "" {
		c.JSON(400, "name can't be none")
		return
	}

	var deployInfo v1.Deployment
	if err := c.BindJSON(&deployInfo); err != nil {
		c.JSON(400, "body error")
		return
	}

	deploy, err := client.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		c.JSON(500, "get deploy error")
		return
	}

	deployInfo.ResourceVersion = deploy.ResourceVersion
	resultDeployment, err := client.AppsV1().Deployments(namespace).Update(&deployInfo)
	if err != nil {
		c.JSON(500, "update deploy error")
		return
	}

	resultDeployment.Kind = "Deployment"
	resultDeployment.APIVersion = "apps/v1"

	c.JSON(200, resultDeployment)
}

func DeleteDeployHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")
	name := c.Query("name")
	if name == "" {
		c.JSON(400, "name can't be none")
		return
	}

	policy := metav1.DeletePropagationForeground
	err := client.AppsV1().Deployments(namespace).Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &policy,
	})
	if err != nil {
		c.JSON(500, "delete deploy error")
		return
	}

	c.JSON(200, "OK")
}
