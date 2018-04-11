package main

import (
	"github.com/hashicorp/consul/api"
)

var (
	fromClient *api.Client
	toClient   *api.Client
)

func main() {
	var err error
	fromClient, err = api.NewClient(&api.Config{
		Address:    "10.200.202.46:8500",
		Datacenter: "jream-j",
	})
	if err != nil {
		panic(err)
	}

	toClient, err = api.NewClient(&api.Config{
		Address: "10.200.150.3:8500",
	})
	if err != nil {
		panic(err)
	}

	kvs, err := from()
	if err != nil {
		panic(err)
	}

	for _, v := range kvs {
		err := set(v.Key, v.Value)
		if err != nil {
			panic(err)
		}
	}
}

func from() (api.KVPairs, error) {
	kv, _, err := fromClient.KV().List("service", nil)
	if err != nil {
		return nil, err
	}

	return kv, nil
}

func set(key string, value []byte) error {
	kv := &api.KVPair{
		Key:   key,
		Value: value,
	}
	_, err := toClient.KV().Put(kv, nil)
	if err != nil {
		return err
	}

	return nil
}
