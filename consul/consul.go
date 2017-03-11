package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := NewClient()
	if err != nil {
		panic(err)
	}
	kv := client.KV()

	p := &api.KVPair{
		Key:   "foo",
		Value: []byte("test"),
	}

	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	pair, _, err := kv.Get("test", nil)
	if err != nil {
		panic(err)
	}
	spew.Dump(pair)
	if pair != nil {
		fmt.Println(string(pair.Value))
	}
}

func NewClient() (*api.Client, error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return nil, err
	}

	return client, nil
}
