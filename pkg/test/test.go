package test

import (
	"bytes"
	"client-go-demo/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http/httptest"
	"testing"
)

func pathPrefix(s string) string {
	return fmt.Sprintf("/v1/%s", s)
}

func mockApi(t *testing.T, method, path string, body interface{}) *httptest.ResponseRecorder {
	var b []byte
	switch body.(type) {
	case string:
		b = []byte(body.(string))
	case []byte:
		b = body.([]byte)
	default:
		var err error
		b, err = json.Marshal(body)
		require.Nil(t, err)
	}

	req := httptest.NewRequest(method, pathPrefix(path), bytes.NewReader(b))
	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	mux := gin.Default()
	conf, err := clientcmd.BuildConfigFromFlags(
		"https://10.6.124.52:16443",
		"/Users/york/go/src/github.com/Fish-pro/client-go-demo/config/52.yaml",
	)

	client, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic("that maybe have problem")
	}

	config.Register(mux, client)

	mux.ServeHTTP(w, req)

	require.True(t, w.Code < 400)

	fmt.Println(w)

	return w
}
