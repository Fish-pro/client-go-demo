package client

import (
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentInterface interface {
	Create(*v1.Deployment) (*v1.Deployment, error)
	Update(*v1.Deployment, string) (*v1.Deployment, error)
	Delete(name string) error
	Get(name string) (*v1.Deployment, error)
	List() (*v1.DeploymentList, error)
}

type DeploymentInterfaceGetter interface {
	Deployments(string) DeploymentInterface
}

type Deployment struct {
	client    *kubernetes.Clientset
	namespace string
}

func newDeployment(client *kubernetes.Clientset, namespace string) *Deployment {
	return &Deployment{client: client, namespace: namespace}
}

func (c *Deployment) Create(deployInfo *v1.Deployment) (*v1.Deployment, error) {

	deploy, err := c.client.AppsV1().Deployments(c.namespace).Create(deployInfo)
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

func (c *Deployment) Update(deployInfo *v1.Deployment, name string) (*v1.Deployment, error) {
	deploy, err := c.client.AppsV1().Deployments(c.namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	deployInfo.ResourceVersion = deploy.ResourceVersion
	resultDeployment, err := c.client.AppsV1().Deployments(c.namespace).Update(deployInfo)
	if err != nil {
		return nil, err
	}

	resultDeployment.Kind = "Deployment"
	resultDeployment.APIVersion = "apps/v1"
	return resultDeployment, nil
}

func (c *Deployment) Delete(name string) error {
	policy := metav1.DeletePropagationForeground
	return c.client.AppsV1().Deployments(c.namespace).Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &policy,
	})
}

func (c *Deployment) Get(name string) (*v1.Deployment, error) {
	return c.client.AppsV1().Deployments(c.namespace).Get(name, metav1.GetOptions{})
}

func (c *Deployment) List() (*v1.DeploymentList, error) {
	deploys, err := c.client.AppsV1().Deployments(c.namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	deploys.Kind = "List"
	deploys.APIVersion = "v1"
	for index := range deploys.Items {
		deploys.Items[index].APIVersion = "apps/v1"
		deploys.Items[index].Kind = "Deployment"
	}
	return deploys, nil
}
