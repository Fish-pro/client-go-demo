package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olivere/elastic.v5"
	"time"
)

type Cmd struct {
	Namespace     string    `json:"namespace"`
	Zone          string    `json:"zone"`
	User          string    `json:"user"`
	Message       string    `json:"message"`
	DateTime      time.Time `json:"dateTime"`
	ContainerName string    `json:"containerName"`
	PodName       string    `json:"podName"`
}

type Params struct {
	Namespace     string
	PodName       string
	ContainerName string
	Message       string
}

func GetCmdParams(c *gin.Context) *Params {
	return &Params{
		Namespace:     c.Query("namespace"),
		PodName:       c.Query("podName"),
		ContainerName: c.Query("containerName"),
		Message:       c.Query("message"),
	}
}

func GetCmdResult(client *elastic.Client, index string, params *Params) (*[]Cmd, error) {

	searchResult, err := client.Search(index).Type("cmd").Do(context.Background())
	if err != nil {
		return nil, err
	}

	res := make([]Cmd, 0, len(searchResult.Hits.Hits))
	if searchResult.Hits.TotalHits > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var c Cmd
			err := json.Unmarshal(*hit.Source, &c)
			if err != nil {
				return nil, fmt.Errorf("parse cmd error")
			}
			res = append(res, c)
		}
	}

	return &res, nil
}
