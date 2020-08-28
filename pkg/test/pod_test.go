package test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestListPodHandler(t *testing.T) {
	mockApi(t, http.MethodGet, "/v1/pods", nil)
}

func TestGetPodHandler(t *testing.T) {
	podName := "myapp-2048-7945b4b584-wvtrf"
	q := url.Values{}
	q.Set("namespace", "default")
	mockApi(t, http.MethodGet, fmt.Sprintf("/v1/pods/%s?%s", podName, q.Encode()), nil)
}

func TestEventsOfPodHandler(t *testing.T) {
	podName := "myapp-2048-7945b4b584-wvtrf"
	q := url.Values{}
	q.Set("namespace", "default")
	mockApi(t, http.MethodGet, fmt.Sprintf("/v1/pods/%s/events?%s", podName, q.Encode()), nil)
}
