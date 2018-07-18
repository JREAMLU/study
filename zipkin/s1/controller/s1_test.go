package controller

import (
	"context"
	"testing"

	proto "github.com/JREAMLU/study/zipkin/s1/proto"
	microClient "github.com/micro/go-plugins/client/grpc"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStart(t *testing.T) {
	serviceName := "go.micro.srv.s1"
	Convey("start", t, func() {
		c := microClient.NewClient()
		client := proto.S1ServiceClient(serviceName, c)
		resp, err := client.AHello(context.Background(), &proto.AHelloRequest{
			Name: "LBJ",
		})
		if err != nil {
			t.Log("err: ", err)
			return
		}

		t.Log("resp: ", resp.Greeting)
	})
}
