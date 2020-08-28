package service

import (
	"client-go-demo/pkg/middleware"
	. "client-go-demo/pkg/util"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

func ListServiceHandler(client *kubernetes.Clientset, c *gin.Context) {
	namespace := c.Query("namespace")
	labelSelector := c.Query("labelSelector")

	serviceList, err := client.CoreV1().Services(namespace).List(metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get services error:%s", err.Error())
		c.JSON(500, "get services error")
		return
	}
	serviceList.Kind = "List"
	serviceList.APIVersion = "v1"

	for index := range serviceList.Items {
		serviceList.Items[index].APIVersion = "v1"
		serviceList.Items[index].Kind = "Service"
	}
	c.JSON(200, serviceList)

}

func GetServiceHandler(client *kubernetes.Clientset, c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Param("name")

	service, err := client.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get service error:%s", err.Error())
		c.JSON(500, "get service error")
		return
	}
	service.APIVersion = "v1"
	service.Kind = "Service"
	c.JSON(200, service)
}

func CreateServiceHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")

	serviceRequest := &v1.Service{}
	err := c.BindJSON(serviceRequest)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get create service body error:%s", err.Error())
		c.JSON(400, "create service body error")
		return
	}

	service, err := client.CoreV1().Services(namespace).Create(serviceRequest)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "create service error:%s", err.Error())
		c.JSON(500, "create service error")
		return
	}
	service.APIVersion = "v1"
	service.Kind = "Service"
	c.JSON(200, service)
}

func UpdateServiceHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")
	name := c.Param("name")

	serviceRequest := &v1.Service{}

	err := c.BindJSON(serviceRequest)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get update service body error:%s", err.Error())
		c.JSON(500, "get update service body error")
		return
	}

	service, err := client.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get service error:%s", err.Error())
		c.JSON(500, "get service error")
		return
	}
	serviceRequest.ResourceVersion = service.ResourceVersion

	updateService, err := client.CoreV1().Services(namespace).Update(serviceRequest)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "update service error:%s", err.Error())
		c.JSON(500, "update service error")
		return
	}

	updateService.APIVersion = "v1"
	updateService.Kind = "Service"
	c.JSON(200, updateService)
}

func DeleteServiceHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")
	name := c.Param("name")

	err := client.CoreV1().Services(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "delete service error:%s", err.Error())
		c.JSON(500, "delete service error")
		return
	}

	c.JSON(200, "ok")
}

func ListEventsOfServiceHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")
	name := c.Param("name")

	service, err := client.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get service error:%s", err.Error())
		c.JSON(500, "get service error")
		return
	}

	eventList, err := client.CoreV1().Events(namespace).Search(scheme.Scheme, service)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get service events error:%s", err.Error())
		c.JSON(500, "get service events error")
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

func ListPodsOfServiceHandler(client *kubernetes.Clientset, c *gin.Context) {

	namespace := c.Query("namespace")
	name := c.Param("name")

	service, err := client.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get service error:%s", err.Error())
		c.JSON(500, "get service error")
		return
	}

	if len(service.Spec.Selector) == 0 {
		var podList v1.PodList
		podList.APIVersion = "v1"
		podList.Kind = "List"
		podList.Items = []v1.Pod{}
		c.JSON(200, podList)
		return
	}

	selector := &metav1.LabelSelector{
		MatchLabels: service.Spec.Selector,
	}
	labelSelector, err := metav1.LabelSelectorAsSelector(selector)
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "parse to selector error:%s", err.Error())
		c.JSON(500, "parse to selector error")
		return
	}

	podList, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{
		LabelSelector: labelSelector.String(),
	})
	if err != nil {
		Logger.Errorf(middleware.GetReqId(c), "get service pods error:%s", err.Error())
		c.JSON(500, "get service pods error")
		return
	}

	podList.APIVersion = "v1"
	podList.Kind = "List"
	for index := range podList.Items {
		podList.Items[index].APIVersion = "v1"
		podList.Items[index].Kind = "Pod"
	}

	c.JSON(200, podList)
}
