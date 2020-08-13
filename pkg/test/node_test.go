package test

import (
	"net/http"
	"testing"
)

func TestListNodeHandler(t *testing.T) {
	mockApi(t, http.MethodGet, "nodes", nil)
}
