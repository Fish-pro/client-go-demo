package config

import (
	"client-go-demo/pkg/app/deployment"
	"client-go-demo/pkg/app/node"
	"client-go-demo/pkg/app/service"
	"client-go-demo/pkg/middleware"
	. "client-go-demo/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

type ServerConfig struct {
	Server   *Server
	Kube     *Kube
	LogLevel string
}

type Server struct {
	Host string
	Port string
}

type Kube struct {
	MasterUrl            string
	KubernetesConfigPath string
}

func (s *Server) GetServerAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

func (k *Kube) New() (*kubernetes.Clientset, error) {
	clusterConfig, err := clientcmd.BuildConfigFromFlags(
		k.MasterUrl,
		k.KubernetesConfigPath,
	)

	if err != nil {
		Logger.Errorf("main", "build cluster config error")
		return nil, err
	}

	client, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		Logger.Errorf("main", "that maybe have problem")
		return nil, err
	}
	return client, nil
}

func getEnvOrDefault(key string, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	} else {
		return def
	}
}

func GetConfigFromENV() *ServerConfig {
	return &ServerConfig{
		Server: &Server{
			Host: getEnvOrDefault("HTTP_HOST", "0.0.0.0"),
			Port: getEnvOrDefault("HTTP_PORT", "9090"),
		},
		Kube: &Kube{
			MasterUrl: getEnvOrDefault("MASTER_URL", "https://10.6.124.52:16443"),
			KubernetesConfigPath: getEnvOrDefault(
				"CLUSTER_CONFIG_PATH",
				"/Users/york/go/src/github.com/Fish-pro/client-go-demo/config/52.yaml",
			),
		},
		LogLevel: getEnvOrDefault("LOG_LEVEL", "DEBUG"),
	}
}

func Register(e *gin.Engine, client *kubernetes.Clientset) {

	e.GET("/ping", func(c *gin.Context) {
		Logger.Debug(middleware.GetReqId(c), "test debug level")
		Logger.Debugf(middleware.GetReqId(c), "test debugf level:%s", "hello")
		Logger.Info(middleware.GetReqId(c), "test debug level")
		Logger.Infof(middleware.GetReqId(c), "test debug level:%s", "hello")
		Logger.Warn(middleware.GetReqId(c), "test debug level")
		Logger.Warnf(middleware.GetReqId(c), "test debug level:%s", "hello")
		Logger.Error(middleware.GetReqId(c), "test debug level")
		Logger.Errorf(middleware.GetReqId(c), "test debug level:%s", "hello")
		c.JSON(200, "pong")
	})

	v1Router := e.Group("/v1")

	// deployment app
	deployment.DeployRouter(client, v1Router)

	// node app
	node.NodeRouter(client, v1Router)

	// service app
	service.SvcRouter(client, v1Router)

}
