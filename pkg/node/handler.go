package node

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/reference"
)

func ListNodeHandler(client *kubernetes.Clientset, c *gin.Context) {

	nodeList, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		c.JSON(400, "failed to get node list")
		return
	}

	nodeList.Kind = "List"
	nodeList.APIVersion = "v1"
	for index := range nodeList.Items {
		nodeList.Items[index].APIVersion = "apps/v1"
		nodeList.Items[index].Kind = "Node"
	}

	c.JSON(200, nodeList)
}

func GetNodeHandler(client *kubernetes.Clientset, c *gin.Context) {

	name := c.Param("name")

	node, err := client.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		c.JSON(500, "get node error")
		return
	}
	node.Kind = "Node"
	node.APIVersion = "apps/v1"

	c.JSON(200, node)
}

func ListPodsOfNodeHandler(client *kubernetes.Clientset, c *gin.Context) {

	name := c.Param("name")

	opts := metav1.ListOptions{FieldSelector: fmt.Sprintf("spec.nodeName=%s", name)}

	podList, err := client.CoreV1().Pods(metav1.NamespaceAll).List(opts)
	if err != nil {
		c.JSON(500, "get node pods error")
		return
	}

	podList.Kind = "List"
	podList.APIVersion = "v1"
	for index := range podList.Items {
		podList.Items[index].APIVersion = "apps/v1"
		podList.Items[index].Kind = "Pod"
	}

	c.JSON(200, podList)
}

func UpdateNodeHandler(client *kubernetes.Clientset, c *gin.Context) {
	name := c.Param("name")

	updateNode := &v1.Node{}
	err := c.BindJSON(updateNode)
	if err != nil {
		c.JSON(500, "get body error")
		return
	}

	node, err := client.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		c.JSON(500, "fail to get node")
		return
	}

	updateNode.ResourceVersion = node.ResourceVersion

	nodeInfo, err := client.CoreV1().Nodes().Update(updateNode)
	if err != nil {
		c.JSON(500, "failed to update node")
		return
	}

	c.JSON(200, nodeInfo)
}

func ListEventsOfNodeHandler(client *kubernetes.Clientset, c *gin.Context) {
	name := c.Param("name")

	node, err := client.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		c.JSON(500, "fetch node error")
		return
	}

	ref, err := reference.GetReference(scheme.Scheme, node)
	if err != nil {
		c.JSON(500, err)
		return
	}
	ref.UID = types.UID(ref.Name)

	eventList, err := client.CoreV1().Events(metav1.NamespaceAll).Search(scheme.Scheme, ref)
	if err != nil {
		c.JSON(500, "failed to get events")
		return
	}

	eventList.APIVersion = "v1"
	eventList.Kind = "List"
	for index := range eventList.Items {
		eventList.Items[index].APIVersion = "v1"
		eventList.Items[index].Kind = "Event"
	}
	c.JSON(200, eventList)
}
