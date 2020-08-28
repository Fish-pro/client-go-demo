package pod

import (
	"client-go-demo/pkg/middleware"
	. "client-go-demo/pkg/util"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

func ListPodHandler(client *kubernetes.Clientset, c *gin.Context) {
	namespace := c.Query("namespace")
	labelSelector := c.Query("labelSelector")

	pods, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get pods error:%s", err.Error())
		c.JSON(500, "get pods error")
		return
	}

	pods.Kind = "List"
	pods.APIVersion = "v1"

	for index := range pods.Items {
		pods.Items[index].APIVersion = "v1"
		pods.Items[index].Kind = "Pod"
	}
	c.JSON(200, pods)
}

func GetPodHandler(client *kubernetes.Clientset, c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")

	pod, err := client.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get pod error:%s", err.Error())
		c.JSON(500, "get pod error")
		return
	}

	pod.APIVersion = "v1"
	pod.Kind = "Pod"
	c.JSON(200, pod)
}

func CreatePodHandler(client *kubernetes.Clientset, c *gin.Context) {
	namespace := c.Query("namespace")

	body := &v1.Pod{}
	err := c.BindJSON(body)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get create pod body error:%s", err.Error())
		c.JSON(400, "create pod body error")
		return
	}

	pod, err := client.CoreV1().Pods(namespace).Create(body)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "create pod error:%s", err.Error())
		c.JSON(500, "create pod error")
		return
	}

	pod.APIVersion = "v1"
	pod.Kind = "Pod"

	c.JSON(200, pod)
}

func UpdatePodHandler(client *kubernetes.Clientset, c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")

	body := &v1.Pod{}
	err := c.BindJSON(body)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get update pod body error:%s", err.Error())
		c.JSON(400, "update pod body error")
		return
	}

	pod, err := client.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get pod error:%s", err.Error())
		c.JSON(500, "get pod error")
		return
	}

	body.ResourceVersion = pod.ResourceVersion

	updatePod, err := client.CoreV1().Pods(namespace).Update(body)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "update pod error:%s", err.Error())
		c.JSON(500, "update pod error")
		return
	}

	updatePod.APIVersion = "v1"
	updatePod.Kind = "Pod"

	c.JSON(200, updatePod)
}

func DeletePodHandler(client *kubernetes.Clientset, c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")

	err := client.CoreV1().Pods(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "delete pod error:%s", err.Error())
		c.JSON(500, "delete pod error")
		return
	}

	c.JSON(200, "ok")
}

func ListEventsOfPodHandler(client *kubernetes.Clientset, c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")

	pod, err := client.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get pod error:%s", err.Error())
		c.JSON(500, "get pod error")
		return
	}

	eventList, err := client.CoreV1().Events(namespace).Search(scheme.Scheme, pod)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get pod events error:%s", err.Error())
		c.JSON(500, "get pod events error")
		return
	}

	eventList.Kind = "Event"
	eventList.APIVersion = "v1"
	for index := range eventList.Items {
		eventList.Items[index].APIVersion = "v1"
		eventList.Items[index].Kind = "Event"
	}

	c.JSON(200, eventList)
}
