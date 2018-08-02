package controller

import (
	"context"
	"testing"

	proto "github.com/JREAMLU/study/zipkin/s2/proto"
	microClient "github.com/micro/go-plugins/client/grpc"
)

func BenchmarkStart(b *testing.B) {
	serviceName := "go.micro.srv.s2"
	c := microClient.NewClient()
	client := proto.S2ServiceClient(serviceName, c)
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.BHello(context.Background(), &proto.BHelloRequest{
			Name: "LBJ",
		})
		if err != nil {
			b.Log("err: ", err)
			return
		}
	}
	b.StopTimer()
}
