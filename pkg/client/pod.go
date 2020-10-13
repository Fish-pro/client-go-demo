package client

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
)

type Pod struct {
	client    *kubernetes.Clientset
	namespace string
}

type PodInterface interface {
	Get(name string) (*v1.Pod, error)
	Create(*v1.Pod) (*v1.Pod, error)
	Update(*v1.Pod) (*v1.Pod, error)
	Delete(string) error
	List() (*v1.PodList, error)
	Events(string) (*v1.EventList, error)
}

type PodInterfaceGetter interface {
	Pods(string) PodInterface
}

func newPod(client *kubernetes.Clientset, namespace string) *Pod {
	return &Pod{client: client, namespace: namespace}
}

func (c *Pod) Get(name string) (*v1.Pod, error) {
	pod, err := c.client.CoreV1().Pods(c.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	pod.APIVersion = "v1"
	pod.Kind = "Pod"
	return pod, nil
}

func (c *Pod) Create(pod *v1.Pod) (*v1.Pod, error) {
	pod, err := c.client.CoreV1().Pods(c.namespace).Create(pod)
	if err != nil {
		return nil, err
	}

	pod.APIVersion = "v1"
	pod.Kind = "Pod"
	return pod, nil
}

func (c *Pod) Update(pod *v1.Pod) (*v1.Pod, error) {
	body, err := c.client.CoreV1().Pods(c.namespace).Get(pod.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	body.ResourceVersion = pod.ResourceVersion

	updatePod, err := c.client.CoreV1().Pods(c.namespace).Update(body)
	if err != nil {
		return nil, err
	}

	updatePod.APIVersion = "v1"
	updatePod.Kind = "Pod"

	return updatePod, nil
}

func (c *Pod) Delete(name string) error {
	return c.client.CoreV1().Pods(c.namespace).Delete(name, &metav1.DeleteOptions{})
}

func (c *Pod) List() (*v1.PodList, error) {
	pods, err := c.client.CoreV1().Pods(c.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	pods.Kind = "List"
	pods.APIVersion = "v1"

	for index := range pods.Items {
		pods.Items[index].APIVersion = "v1"
		pods.Items[index].Kind = "Pod"
	}
	return pods, nil
}

func (c *Pod) Events(name string) (*v1.EventList, error) {
	pod, err := c.client.CoreV1().Pods(c.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	eventList, err := c.client.CoreV1().Events(c.namespace).Search(scheme.Scheme, pod)
	if err != nil {
		return nil, err
	}

	eventList.Kind = "Event"
	eventList.APIVersion = "v1"
	for index := range eventList.Items {
		eventList.Items[index].APIVersion = "v1"
		eventList.Items[index].Kind = "Event"
	}
	return eventList, nil
}
