package es

import (
	"gopkg.in/olivere/elastic.v5"
)

const (
	EsUrl   = "http://47.99.56.228:9200"
	EsToken = ""
)

type Es struct {
	Url   string
	Token string
}

func New() (*elastic.Client, error) {

	es := &Es{
		Url:   EsUrl,
		Token: EsToken,
	}

	client, err := elastic.NewClient(elastic.SetURL(es.Url), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	return client, nil
}
