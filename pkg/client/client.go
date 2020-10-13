package client

import "k8s.io/client-go/kubernetes"

type Client struct {
	client *kubernetes.Clientset
}

type MainInterface interface {
	NodeInterfaceGetter
	DeploymentInterfaceGetter
	PodInterfaceGetter
	ServiceInterfaceGetter
}

var _ MainInterface = &Client{}

func New(client *kubernetes.Clientset) *Client {
	return &Client{client: client}
}

// node client
func (c *Client) Nodes() NodeInterface {
	return newNodes(c.client)
}

// deployment client
func (c *Client) Deployments(namespace string) DeploymentInterface {
	return newDeployment(c.client, namespace)
}

// pod client
func (c *Client) Pods(namespace string) PodInterface {
	return newPod(c.client, namespace)
}

// service client
func (c *Client) Services(namespace string) ServiceInterface {
	return newService(c.client, namespace)
}
