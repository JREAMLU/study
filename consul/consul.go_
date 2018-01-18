package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/consul/api"
)

func main() {
	fmt.Println(123)
	client, err := NewClient()
	if err != nil {
		panic(err)
	}
	kv := client.KV()

	// p := &api.KVPair{
	// 	Key:   "foo",
	// 	Value: []byte("test"),
	// }
	//
	// _, err = kv.Put(p, nil)
	// if err != nil {
	// 	panic(err)
	// }

	pair, _, err := kv.Get("mysql", nil)
	if err != nil {
		panic(err)
	}
	spew.Dump(pair)
	if pair != nil {
		fmt.Println(string(pair.Value))
	}

	agent := client.Agent()
	reg := &api.AgentServiceRegistration{
		Name: "api",
		Tags: []string{"bar", "baz"},
		Port: 8000,
		Check: &api.AgentServiceCheck{
			TTL: "15s",
		},
	}
	err = agent.ServiceRegister(reg)
	if err != nil {
		fmt.Println(err)
	}

	err = agent.ServiceDeregister("foo")
	if err != nil {
		fmt.Println(err)
	}

	services, err := agent.Services()
	fmt.Println("++++++++++++: ", services["api"].Service)
}

func NewClient() (*api.Client, error) {
	config := &api.Config{
		Address: "10.211.55.5:8500",
	}
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
