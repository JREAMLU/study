package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gopkg.in/olivere/elastic.v3"
)

type Member struct {
	User     string `json:"user"`
	Nickname string `json:"nickname"`
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("panic:", x)
		}
	}()

	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}

	info, code, err := client.Ping("http://127.0.0.1:9200").Do()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s \n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s \n", esversion)

	exists, err := client.IndexExists("member").Do()
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := client.CreateIndex("member").Do()
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			fmt.Println("Not acknowledged: ", createIndex.Acknowledged)
		}
	}

	/*
			member1 := Member{User: "jream", Nickname: "fjh"}
			put1, err := client.Index().
				Index("member").
				Type("family").
				Id("1").
				BodyJson(member1).
				Do()
			if err != nil {
				// Handle error
				panic(err)
			}

			member2 := Member{User: "jream", Nickname: "fjh34D"}
			put2, err := client.Index().
				Index("member").
				Type("family").
				Id("2").
				BodyJson(member2).
				Do()
			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
			fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put2.Index, put2.Type)

		get1, err := client.Get().
			Index("member").
			Type("family").
			Id("1").
			Do()
		if err != nil {
			// Handle error
			panic(err)
		}
		if get1.Found {
			fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		}
	*/

	termQuery := elastic.NewTermQuery("user", "jream")
	searchResult, err := client.Search().
		Index("member").    // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do()                // execute
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var mem Member
	for _, item := range searchResult.Each(reflect.TypeOf(mem)) {
		m := item.(Member)
		fmt.Printf("Member by %s: %s\n", m.User, m.Nickname)
	}

	fmt.Printf("Found a total of %d family\n", searchResult.TotalHits())

	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d family\n", searchResult.Hits.TotalHits)

		for _, hit := range searchResult.Hits.Hits {
			var m Member
			err := json.Unmarshal(*hit.Source, &m)
			if err != nil {
			}

			fmt.Printf("Member by %s: %s\n", m.User, m.Nickname)
		}
	} else {
		fmt.Print("Found no family\n")
	}

	script := elastic.NewScript("ctx._source.nickname = _nickname").Param("_nickname", "fjh301")
	update, err := client.Update().Index("member").Type("family").Id("1").
		Script(script).
		Upsert(map[string]interface{}{"nickname": "fjh"}).
		Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("New version of member %q is now %d", update.Id, update.Version)

	_, err = client.Flush().Index("member").Do()
	if err != nil {
		panic(err)
	}

	/*
		deleteIndex, err := client.DeleteIndex("member").Do()
		if err != nil {
			panic(err)
		}
		if !deleteIndex.Acknowledged {
			fmt.Println("Not acknowledged: ", deleteIndex.Acknowledged)
		}
		/*/
}
