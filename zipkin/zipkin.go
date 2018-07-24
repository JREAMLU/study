package main

import (
	"context"
	"fmt"
	"os"

	"github.com/micro/go-micro/metadata"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
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
	// sc, _ := span.Context().(zipkin.SpanContext)
	// fmt.Println("++++++++++++: ", sc.TraceID.ToHex())
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	server1(ctx, tracer)
	server2(ctx, tracer)
	server3(ctx, tracer)
}

func server1(ctx context.Context, t opentracing.Tracer) {
	span, _ := opentracing.StartSpanFromContext(ctx, "call server1")
	span.SetTag("host", "8888")
	span.LogFields(
		log.String("type", "fields"),
		log.String("error", "mis match"),
	)
	span.LogKV("name", "luj", "skill", "go", "ab", 99)
	span.LogEvent("call server1 over")
	span.Finish()
}

func server2(ctx context.Context, t opentracing.Tracer) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}

	span, _ := opentracing.StartSpanFromContext(ctx, "call server2")
	err := span.Tracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	span.LogEvent("server2")
	span.Finish()
}

func server3(ctx context.Context, t opentracing.Tracer) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}

	var span opentracing.Span
	wireContext, err := t.Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	if err != nil {
		fmt.Println("++++++++++++: ", err)
		span = t.StartSpan("server3", opentracing.ChildOf(wireContext))
	} else {
		span = t.StartSpan("server3")
	}

	err = span.Tracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
	if err != nil {
		fmt.Println("err2: ", err)
		return
	}

	// ctx = opentracing.ContextWithSpan(ctx, span)
	// ctx = metadata.NewContext(ctx, md)

	span.LogEvent("server3")
	span.Finish()
}
