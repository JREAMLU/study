package main

import (
	"context"
	"log"

	"github.com/JREAMLU/study/micro/service/proto"
	"github.com/micro/go-plugins/client/grpc"
)

// grpc "github.com/micro/go-grpc"
// "github.com/micro/go-micro/metadata"

func main() {
	/*
		service := grpc.NewService(
			micro.Name("greeter.client"),
		)

		greeter := proto.NewGreeterClient("greeter", service.Client())

		rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "JREAM"})
		if err != nil {
			log.Println(err)
		}

		log.Println(rsp.Nickname)
	*/

	/*
		service := grpc.NewService()
		service.Init()

		greeter := proto.NewGreeterClient("greeter", service.Client())

		ctx := metadata.NewContext(context.Background(), map[string]string{
			"X-User-Id": "lu",
			"X-From-Id": "script",
		})

		rsp, err := greeter.Hello(ctx, &proto.HelloRequest{Name: "JREAM"})
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(rsp.Nickname)
	*/

	// resp, err := hello()
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("result: ", resp.Nickname)

	resp, err := hello()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result: ", resp.Nickname)
}

// func hello() (*proto.HelloResponse, error) {
// 	req := client.NewRequest("greeter", "Greeter.Hello", &proto.HelloRequest{Name: "JREAM"})
// 	resp := &proto.HelloResponse{}
// 	err := client.Call(context.TODO(), req, resp)
//
// if err != nil {
// 	log.Println(err)
// 	return nil, err
// }
//
// return resp, nil
// }

func hello() (*proto.HelloResponse, error) {
	client := grpc.NewClient()
	req := client.NewRequest("greeter", "Greeter.Hello", &proto.HelloRequest{Name: "JREAM"})
	resp := &proto.HelloResponse{}
	err := client.Call(context.TODO(), req, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
