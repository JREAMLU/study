package controller

import (
	"context"
	"sync"
	"testing"

	proto "github.com/JREAMLU/study/zipkin/s1/proto"
	microClient "github.com/micro/go-plugins/client/grpc"

	jopentracing "github.com/JREAMLU/j-kit/go-micro/trace/opentracing"
	"github.com/micro/go-micro/metadata"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStart(t *testing.T) {
	serviceName := "go.micro.srv.s1"

	// add trace toggle
	md := make(map[string]string)
	md[jopentracing.ZipkinToggle] = "1"
	ctx := metadata.NewContext(context.Background(), md)

	Convey("start", t, func() {
		c := microClient.NewClient()
		client := proto.S1ServiceClient(serviceName, c)
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				resp, err := client.AHello(ctx, &proto.AHelloRequest{
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

// func TestStartB(t *testing.T) {
// 	serviceName := "go.micro.srv.s1"
// 	Convey("start", t, func() {
// 		c := microClient.NewClient()
// 		client := proto.S1ServiceClient(serviceName, c)
// 		for i := 0; i < 10; i++ {
// 			resp, err := client.AHello(context.Background(), &proto.AHelloRequest{
// 				Name: "LBJ",
// 			})
// 			if err != nil {
// 				t.Log("err: ", err)
// 				continue
// 			}
//
// 			t.Log("resp: ", resp.Greeting)
// 		}
// 	})
// }

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
