package config

import (
	"client-go-demo/pkg/app/deployment"
	"client-go-demo/pkg/app/node"
	"client-go-demo/pkg/middleware"
	"client-go-demo/pkg/util"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func Register(r *gin.Engine, client *kubernetes.Clientset) {

	r.GET("/ping", func(c *gin.Context) {
		util.Logger.Debug(middleware.GetReqId(c), "test debug level")
		util.Logger.Debugf(middleware.GetReqId(c), "test debugf level:%s", "hello")
		util.Logger.Info(middleware.GetReqId(c), "test debug level")
		util.Logger.Infof(middleware.GetReqId(c), "test debug level:%s", "hello")
		util.Logger.Warn(middleware.GetReqId(c), "test debug level")
		util.Logger.Warnf(middleware.GetReqId(c), "test debug level:%s", "hello")
		util.Logger.Error(middleware.GetReqId(c), "test debug level")
		util.Logger.Errorf(middleware.GetReqId(c), "test debug level:%s", "hello")
		c.JSON(200, "pong")
	})

	v1Router := r.Group("/v1")

	// deployment app
	deployment.DeployRouter(client, v1Router)

	// node app
	node.NodeRouter(client, v1Router)

}
