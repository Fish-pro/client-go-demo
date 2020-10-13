package client

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

type Service struct {
	client    *kubernetes.Clientset
	namespace string
}

type ServiceInterface interface {
	List() (*v1.ServiceList, error)
	Get(string) (*v1.Service, error)
	Create(*v1.Service) (*v1.Service, error)
	Update(*v1.Service, string) (*v1.Service, error)
	Delete(string) error
	Events(string) (*v1.EventList, error)
	Pods(string) (*v1.PodList, error)
}

type ServiceInterfaceGetter interface {
	Services(string) ServiceInterface
}

func newService(client *kubernetes.Clientset, namespace string) *Service {
	return &Service{client: client, namespace: namespace}
}

func (c *Service) List() (*v1.ServiceList, error) {
	serviceList, err := c.client.CoreV1().Services(c.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	serviceList.Kind = "List"
	serviceList.APIVersion = "v1"

	for index := range serviceList.Items {
		serviceList.Items[index].APIVersion = "v1"
		serviceList.Items[index].Kind = "Service"
	}
	return serviceList, nil
}

func (c *Service) Get(name string) (*v1.Service, error) {
	service, err := c.client.CoreV1().Services(c.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	service.APIVersion = "v1"
	service.Kind = "Service"
	return service, nil
}

func (c *Service) Create(serviceInfo *v1.Service) (*v1.Service, error) {

	service, err := c.client.CoreV1().Services(c.namespace).Create(serviceInfo)
	if err != nil {
		return nil, err
	}
	service.APIVersion = "v1"
	service.Kind = "Service"
	return service, nil
}

func (c *Service) Update(serviceInfo *v1.Service, name string) (*v1.Service, error) {
	service, err := c.client.CoreV1().Services(c.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	serviceInfo.ResourceVersion = service.ResourceVersion

	updateService, err := c.client.CoreV1().Services(c.namespace).Update(serviceInfo)
	if err != nil {
		return nil, err
	}

	updateService.APIVersion = "v1"
	updateService.Kind = "Service"

	return updateService, nil
}

func (c *Service) Delete(name string) error {
	return c.client.CoreV1().Services(c.namespace).Delete(name, &metav1.DeleteOptions{})
}

func (c *Service) Events(name string) (*v1.EventList, error) {
	service, err := c.client.CoreV1().Services(c.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	eventList, err := c.client.CoreV1().Events(c.namespace).Search(scheme.Scheme, service)
	if err != nil {
		return nil, err
	}
	eventList.APIVersion = "v1"
	eventList.Kind = "List"
	for index := range eventList.Items {
		eventList.Items[index].APIVersion = "v1"
		eventList.Items[index].Kind = "Event"
	}
	return eventList, nil
}

func (c *Service) Pods(name string) (*v1.PodList, error) {

	service, err := c.client.CoreV1().Services(c.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if len(service.Spec.Selector) == 0 {
		var podList v1.PodList
		podList.APIVersion = "v1"
		podList.Kind = "List"
		podList.Items = []v1.Pod{}
		return &podList, nil
	}

	selector := &metav1.LabelSelector{
		MatchLabels: service.Spec.Selector,
	}
	labelSelector, err := metav1.LabelSelectorAsSelector(selector)
	if err != nil {
		return nil, err
	}

	podList, err := c.client.CoreV1().Pods(c.namespace).List(metav1.ListOptions{
		LabelSelector: labelSelector.String(),
	})
	if err != nil {
		return nil, err
	}

	podList.APIVersion = "v1"
	podList.Kind = "List"
	for index := range podList.Items {
		podList.Items[index].APIVersion = "v1"
		podList.Items[index].Kind = "Pod"
	}
	return podList, nil
}
