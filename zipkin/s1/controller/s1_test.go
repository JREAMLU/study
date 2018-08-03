package controller

import (
	"context"
	"sync"
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
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				resp, err := client.AHello(context.Background(), &proto.AHelloRequest{
					Name: "LBJ",
				})
				if err != nil {
					t.Log("err: ", err)
					return
				}

				t.Log("resp: ", resp.Greeting)
			}()
		}
		wg.Wait()
	})
}

func BenchmarkStart(b *testing.B) {
	serviceName := "go.micro.srv.s1"
	c := microClient.NewClient()
	client := proto.S1ServiceClient(serviceName, c)
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.AHello(context.Background(), &proto.AHelloRequest{
			Name: "LBJ",
		})
		if err != nil {
			b.Log("err: ", err)
			return
		}

		// b.Log("resp: ", resp.Greeting)
	}
	b.StopTimer()
}
