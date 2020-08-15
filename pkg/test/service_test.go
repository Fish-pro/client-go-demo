package test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestListServiceHandler(t *testing.T) {
	mockApi(t, http.MethodGet, "services", nil)
}

func TestGetServiceHandler(t *testing.T) {
	serviceName := "nginx307-v1"
	q := url.Values{}
	q.Set("namespace", "testproject")
	mockApi(t, http.MethodGet, fmt.Sprintf("services/%s?%s", serviceName, q.Encode()), nil)
}

func TestListEventsOfServiceHandler(t *testing.T) {
	serviceName := "nginx307-v1"
	q := url.Values{}
	q.Set("namespace", "testproject")
	mockApi(t, http.MethodGet, fmt.Sprintf("services/%s/events?%s", serviceName, q.Encode()), nil)
}

func TestListPodsOfServiceHandler(t *testing.T) {
	serviceName := "nginx307-v1"
	q := url.Values{}
	q.Set("namespace", "testproject")
	mockApi(t, http.MethodGet, fmt.Sprintf("services/%s/pods?%s", serviceName, q.Encode()), nil)
}
