package test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestListDeployHandler(t *testing.T) {
	mockApi(t, http.MethodGet, "deployments", nil)
}

func TestGetDeployHandler(t *testing.T) {
	deployName := "nginx75-v1"
	q := url.Values{}
	q.Set("namespace", "testproject")

	mockApi(t, http.MethodGet, fmt.Sprintf("deployments/%s?%s", deployName, q.Encode()), nil)
}
