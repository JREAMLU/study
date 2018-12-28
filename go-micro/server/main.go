package main

import (
	"fmt"

	"github.com/JREAMLU/study/go-micro/server/controller"
	pb "github.com/JREAMLU/study/go-micro/server/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	pb.RegisterGreeterHandler(service.Server(), controller.NewGreeterHandler())

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
