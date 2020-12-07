package main

import (
	fmt "fmt"
	"google.golang.org/grpc/metadata"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/opentracing/opentracing-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

var (
	rpcCli ApiServiceClient
)

func init() {
	InitTracer()
	InitRpcClient()
}

func InitTracer() {
	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.64.2:6831",
		},

		ServiceName: "opentrace",
	}

	tracer, _, err := cfg.NewTracer(jaegerConfig.Logger(log.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	fmt.Println("SetGlobalTracer OK")
}

func mustDial(addr string) *grpc.ClientConn {
	c, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*200)),
		grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: time.Second * 3, Backoff: backoff.DefaultConfig}))
	if err != nil {
		panic(err)
	}
	return c
}

func InitRpcClient() {
	rpcCli = NewApiServiceClient(mustDial("localhost:1300"))
}

func waitInter() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-c
}

type MDReaderWriter struct {
	metadata.MD
}

func (c MDReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range c.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c MDReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	c.MD[key] = append(c.MD[key], val)
}

func main() {
	go httpRun()
	go rpcRun()
	waitInter()
}
