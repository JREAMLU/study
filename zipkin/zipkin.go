package main

import (
	"context"
	"fmt"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

var (
	addrs         = []string{"10.200.119.128:9092", "10.200.119.129:9092", "10.200.119.130:9092"}
	hostPort      = "127.0.0.1:8080"
	serviceName   = "go.micro.srv.jxx"
	debug         = true
	sameSpan      = true
	traceID128Bit = true
	kafkaTopic    = "web_log_get"
)

func main() {
	collector, err := zipkin.NewKafkaCollector(
		addrs,
		zipkin.KafkaTopic(kafkaTopic),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin KAFKA collector: %+v\n", err)
		os.Exit(-1)
	}
	defer collector.Close()

	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(sameSpan),
		zipkin.TraceID128Bit(traceID128Bit),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		os.Exit(-1)
	}

	opentracing.InitGlobalTracer(tracer)

	span := opentracing.StartSpan("Run")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	span.LogEvent("log Call server1")
	server1(ctx)
	span.LogEvent("log Call server1 over")
}

func server1(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "call server1")
	span.LogEvent("call server1 over")
	span.Finish()
}
