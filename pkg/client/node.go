package client

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

// NodesGetter has a method to return a NodeInterface.
type NodesGetter interface {
	Nodes(client *kubernetes.Clientset, c *gin.Context) NodeInterface
}

// node interface
type NodeInterface interface {
	List() (*v1.NodeList, error)
	Get(name string) (*v1.Node, error)
	Update(name string, node *v1.Node) (*v1.Node, error)
	Pods(name string) (*v1.PodList, error)
	Events(name string) (*v1.EventList, error)
}

type NodeInterfaceGetter interface {
	Nodes() NodeInterface
}

// struct for interface
type Nodes struct {
	client *kubernetes.Clientset
}

// new node struct
func newNodes(client *kubernetes.Clientset) *Nodes {
	return &Nodes{client: client}
}

func (n *Nodes) List() (*v1.NodeList, error) {
	nodeList, err := n.client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return nodeList, err
	}

	nodeList.Kind = "List"
	nodeList.APIVersion = "v1"
	for index := range nodeList.Items {
		nodeList.Items[index].APIVersion = "apps/v1"
		nodeList.Items[index].Kind = "Node"
	}
	return nodeList, nil
}

func (n *Nodes) Get(name string) (*v1.Node, error) {
	node, err := n.client.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		return node, err
	}
	node.Kind = "Node"
	node.APIVersion = "apps/v1"
	return node, nil
}

func (n *Nodes) Update(name string, updateNode *v1.Node) (*v1.Node, error) {
	node, err := n.client.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	updateNode.ResourceVersion = node.ResourceVersion

	nodeInfo, err := n.client.CoreV1().Nodes().Update(updateNode)
	if err != nil {
		return nil, err
	}
	return nodeInfo, nil
}

func (n *Nodes) Pods(name string) (*v1.PodList, error) {
	opts := metav1.ListOptions{FieldSelector: fmt.Sprintf("spec.nodeName=%s", name)}

	podList, err := n.client.CoreV1().Pods(metav1.NamespaceAll).List(opts)
	if err != nil {
		return nil, err
	}

	podList.Kind = "List"
	podList.APIVersion = "v1"
	for index := range podList.Items {
		podList.Items[index].APIVersion = "apps/v1"
		podList.Items[index].Kind = "Pod"
	}
	return podList, nil
}

func (n *Nodes) Events(name string) (*v1.EventList, error) {
	node, err := n.client.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	ref, err := reference.GetReference(scheme.Scheme, node)
	if err != nil {
		return nil, err
	}
	ref.UID = types.UID(ref.Name)

	eventList, err := n.client.CoreV1().Events(metav1.NamespaceAll).Search(scheme.Scheme, ref)
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
