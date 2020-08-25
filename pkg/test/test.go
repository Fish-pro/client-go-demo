package test

import (
	"bytes"
	"client-go-demo/config"
	"client-go-demo/pkg/middleware"
	"client-go-demo/pkg/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http/httptest"
	"testing"
)

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

	req := httptest.NewRequest(method, path, bytes.NewReader(b))
	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	engine := middleware.SetupGin()
	conf, err := clientcmd.BuildConfigFromFlags(
		"https://10.6.124.55:16443",
		"/Users/york/go/src/github.com/Fish-pro/client-go-demo/config/55.yaml",
	)

	util.Logger.SetLogLevel("WARN")

	client, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic("that maybe have problem")
	}

	config.Register(engine, client)

	engine.ServeHTTP(w, req)

	require.True(t, w.Code < 400)

	// fmt.Println(w)

	return w
}
