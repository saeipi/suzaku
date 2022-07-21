package es

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"suzaku/pkg/common/config"
)

var (
	Client *esClient
)

type esClient struct {
	client *elasticsearch.Client
}

func init() {
	var (
		cfg    elasticsearch.Config
		client *elasticsearch.Client
		err    error
	)
	cfg = elasticsearch.Config{
		Addresses: []string{
			config.Config.Es.Addr,
		},
		Username: config.Config.Es.Username,
		Password: config.Config.Es.Password,
	}
	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	Client = &esClient{client: client}
}
