package util

import (
	"fmt"
	"net"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-plugins/wrapper/trace/opentracing"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

// SetZipkin set zipkin trace
func SetZipkin(service micro.Service, kafkaAddrs []string, kafkaTopic string, hostPort ...string) {
	opts := service.Options()
	setZipkin(opts.Server.Options().Name, opts.Server.Options().Version, kafkaAddrs, kafkaTopic, hostPort...)(&opts)
}

func setZipkin(serviceName, version string, kafkaAddrs []string, kafkaTopic string, hostPort ...string) micro.Option {
	return func(opt *micro.Options) {
		collector, err := zipkin.NewKafkaCollector(
			kafkaAddrs,
			zipkin.KafkaTopic(kafkaTopic),
		)
		if err != nil {
			panic(err)
		}

		var ipAndPort string
		if len(hostPort) == 0 {
			ipAndPort, err = extractAddress("")
		} else {
			ipAndPort, err = extractAddress(hostPort[0])
		}

		if err != nil {
			panic(err)
		}

		recorder := zipkin.NewRecorder(collector, false, ipAndPort, serviceName)
		tracer, err := zipkin.NewTracer(
			recorder,
			zipkin.ClientServerSameSpan(true),
			zipkin.TraceID128Bit(true),
		)
		if err != nil {
			panic(err)
		}

		clientWrap := opentracing.NewClientWrapper(tracer)
		serverWrap := opentracing.NewHandlerWrapper(tracer)

		micro.WrapClient(clientWrap)(opt)
		micro.WrapHandler(serverWrap)(opt)
	}
}

func extractAddress(addr string) (string, error) {
	if len(addr) > 0 && (addr != "0.0.0.0" && addr != "[::]") {
		return addr, nil
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", fmt.Errorf("Failed to get interface addresses! Err: %v", err)
	}

	var ipAddr []byte

	for _, rawAddr := range addrs {
		var ip net.IP
		switch addr := rawAddr.(type) {
		case *net.IPAddr:
			ip = addr.IP
		case *net.IPNet:
			ip = addr.IP
		default:
			continue
		}

		if ip.To4() == nil {
			continue
		}

		if !isPrivateIP(ip.String()) {
			continue
		}

		ipAddr = ip
		break
	}

	if ipAddr == nil {
		return "", fmt.Errorf("No private IP address found, and explicit IP not provided")
	}

	return net.IP(ipAddr).String(), nil
}

func isPrivateIP(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	for _, priv := range privateBlocks {
		if priv.Contains(ip) {
			return true
		}
	}
	return false
}

var (
	privateBlocks []*net.IPNet
)

func init() {
	for _, b := range []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"} {
		if _, block, err := net.ParseCIDR(b); err == nil {
			privateBlocks = append(privateBlocks, block)
		}
	}
}
