package cmd

import (
	"github.com/gin-gonic/gin"
)

func CmdRouter(c *gin.RouterGroup) {
	dRouter := c.Group("/cmd")
	{
		dRouter.GET("", func(c *gin.Context) {
			GetCmdHandler(c)
		})
		dRouter.POST("", func(c *gin.Context) {
			CreateCmdHandler(c)
		})
	}
}
