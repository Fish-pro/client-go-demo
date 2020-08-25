package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestListNodeHandler(t *testing.T) {
	mockApi(t, http.MethodGet, "/v1/nodes", nil)
}

func TestGetNodeHandler(t *testing.T) {
	nodeName := "dce-10-6-124-55"
	mockApi(t, http.MethodGet, fmt.Sprintf("/v1/nodes/%s", nodeName), nil)
}

func TestListPodsOfNodeHandler(t *testing.T) {
	nodeName := "dce-10-6-124-55"
	mockApi(t, http.MethodGet, fmt.Sprintf("/v1/nodes/%s/pods", nodeName), nil)
}

func TestListEventsOfNodeHandler(t *testing.T) {
	nodeName := "dce-10-6-124-55"
	mockApi(t, http.MethodGet, fmt.Sprintf("/v1/nodes/%s/events", nodeName), nil)
}
