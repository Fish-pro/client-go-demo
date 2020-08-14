package middleware

import (
	"client-go-demo/pkg/util"
	"github.com/gin-gonic/gin"
)

const ReqIdKey = "REQIDKEY"
const DefaultReqId = "NoReqId"

// 返回的 header 中的本次请求的流水号，同一个请求中，与上面的 reqid 一致
const XRequestIdHeaderKey = "X-Request-Id"

func ReqIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqId := util.GenerateDataId()
		c.Set(ReqIdKey, reqId)
		c.Writer.Header().Set(XRequestIdHeaderKey, reqId)
		c.Next()
	}
}

func GetReqId(c *gin.Context) string {
	reqId, ok := c.Get(ReqIdKey)
	if !ok {
		// 必须panic，因为属于程序错误，忘记配置 reqid 了
		panic("忘记配置reqid，请检查程序中间件")
	}
	return reqId.(string)
}
