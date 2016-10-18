package core

import (
	"fmt"

	"gopkg.in/olivere/elastic.v3"
)

var ESClient *elastic.Client

func InitElastic(url string) error {
	var err error
	ESClient, err = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return err
	}

	info, code, err := ESClient.Ping(url).Do()
	if err != nil {
		return err
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s \n", code, info.Version.Number)

	esversion, err := ESClient.ElasticsearchVersion(url)
	if err != nil {
		return err
	}
	fmt.Printf("Elasticsearch version %s \n", esversion)

	return nil
}

func CreateIndexElastic(index string) error {
	exists, err := ESClient.IndexExists(index).Do()
	if err != nil {
		return err
	}
	if !exists {
		createIndex, err := ESClient.CreateIndex(index).Do()
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			fmt.Println("Not acknowledged: ", createIndex.Acknowledged)
		}
	}
	return nil
}
