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

	services, meta, err := client.Catalog().Services(&api.QueryOptions{})
	fmt.Println("++++++++++++: ", services, meta, err)
}
