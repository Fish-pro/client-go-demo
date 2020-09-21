package cmd

import (
	"client-go-demo/pkg/client/es"
	"context"
	"github.com/gin-gonic/gin"

	. "client-go-demo/pkg/util"
)

type CreateCmdRequest struct {
	Index string `json:"index"`
	Cmd   *Cmd   `json:"cmd"`
}

func CreateCmdHandler(c *gin.Context) {

	body := &CreateCmdRequest{}
	err := c.BindJSON(body)
	if err != nil {
		Logger.Error("get body error:%s", err.Error())
		c.JSON(400, "params error")
		return
	}

	client, err := es.New()
	if err != nil {
		Logger.Error("get es client error:%s", err.Error())
		c.JSON(500, "get es client error")
		return
	}

	ctx := context.Background()

	exists, err := client.IndexExists(body.Index).Do(ctx)
	if err != nil {
		Logger.Error("check index error:%s", err.Error())
		c.JSON(500, "check index error")
		return
	}

	if !exists {
		_, err := client.CreateIndex(body.Index).Do(ctx)
		if err != nil {
			Logger.Error("create index error:%s", err.Error())
			c.JSON(500, "create index error")
			return
		}
	}

	_, err = client.Index().Index(body.Index).Type("cmd").BodyJson(body.Cmd).Do(ctx)
	if err != nil {
		Logger.Error("failed to write data to es::%s", err.Error())
		c.JSON(500, "failed to write data to es")
		return
	}

	c.JSON(200, "ok")

}

func GetCmdHandler(c *gin.Context) {

	params := GetCmdParams(c)
	index := c.Query("index")

	client, err := es.New()
	if err != nil {
		Logger.Error("get es client error:%s", err.Error())
		c.JSON(500, "get es client error")
		return
	}

	result, err := GetCmdResult(client, index, params)
	if err != nil {
		Logger.Error("query data error:%s", err.Error())
		c.JSON(500, "query data error")
		return
	}

	c.JSON(200, result)

}
