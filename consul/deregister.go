package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := NewClient()
	if err != nil {
		panic(err)
	}

	agent := client.Agent()

	err = agent.ServiceDeregister("http-api")
	if err != nil {
		fmt.Println(err)
	}
}

// NewClient new client
func NewClient() (*api.Client, error) {
	config := &api.Config{
		Address: "10.200.202.77:8500",
	}
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
