package test

import (
	"net/http"
	"testing"
)

func TestListDemoHandler(t *testing.T) {
	mockApi(t, http.MethodGet, "/ping", nil)
}
