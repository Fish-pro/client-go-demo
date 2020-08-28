package test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestListDeployHandler(t *testing.T) {
	mockApi(t, http.MethodGet, "/v1/deployments", nil)
}

func TestGetDeployHandler(t *testing.T) {
	deployName := "dao-2048-dao-2048"
	q := url.Values{}
	q.Set("namespace", "default")

	mockApi(t, http.MethodGet, fmt.Sprintf("/v1/deployments/%s?%s", deployName, q.Encode()), nil)
}
