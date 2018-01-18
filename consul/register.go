package main

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := NewClient()
	if err != nil {
		panic(err)
	}

	agent := client.Agent()
	reg := &api.AgentServiceRegistration{
		Name: "http-api",
		Tags: []string{"v1", "v2"},
		Port: 8001,
		Check: &api.AgentServiceCheck{
			// TTL:      "5s",
			HTTP:     "http://10.200.202.93:8001/status",
			Interval: "5s",
			Timeout:  "1s",
		},
	}
	err = agent.ServiceRegister(reg)
	if err != nil {
		fmt.Println(err)
	}

	// err = agent.ServiceDeregister("http-api")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	services, err := agent.Services()
	if err != nil {
		panic(err)
	}
	fmt.Println("++++++++++++: ", services["http-api"].Service)

	http.HandleFunc("/status", StatusHandler)
	err = http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}
	select {}
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("check status.")
	fmt.Fprint(w, "status ok!")
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
