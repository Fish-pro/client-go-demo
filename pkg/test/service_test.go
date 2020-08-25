package test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestListServiceHandler(t *testing.T) {
	mockApi(t, http.MethodGet, "/v1/services", nil)
}

func TestGetServiceHandler(t *testing.T) {
	serviceName := "helloworld-qx95m-metrics"
	q := url.Values{}
	q.Set("namespace", "default")
	mockApi(t, http.MethodGet, fmt.Sprintf("/v1/services/%s?%s", serviceName, q.Encode()), nil)
}

func TestListEventsOfServiceHandler(t *testing.T) {
	serviceName := "helloworld-qx95m-metrics"
	q := url.Values{}
	q.Set("namespace", "default")
	mockApi(t, http.MethodGet, fmt.Sprintf("/v1/services/%s/events?%s", serviceName, q.Encode()), nil)
}

func TestListPodsOfServiceHandler(t *testing.T) {
	serviceName := "helloworld-qx95m-metrics"
	q := url.Values{}
	q.Set("namespace", "default")
	mockApi(t, http.MethodGet, fmt.Sprintf("/v1/services/%s/pods?%s", serviceName, q.Encode()), nil)
}
