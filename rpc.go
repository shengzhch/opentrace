package main

import (
	"context"
	"net"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

type welcome struct {
}

func (w *welcome) Welcome(ctx context.Context, req *WelcomeReq) (*WelcomeResp, error) {

	s1, ctx := opentracing.StartSpanFromContext(ctx, "welcomedosomething")
	defer s1.Finish()

	//rpcsomething
	func() {
		time.Sleep(time.Millisecond * 200)
	}()

	return &WelcomeResp{}, nil
}

func rpcRun() {
	s := grpc.NewServer(TraceServerOption())
	RegisterApiServiceServer(s, &welcome{})
	lis, err := net.Listen("tcp", "0.0.0.0:1300")
	if err != nil {
		panic(err)
	}
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func TraceServerOption() grpc.ServerOption {
	usi := serverInterceptor()
	return grpc.UnaryInterceptor(usi)
}

func serverInterceptor() grpc.UnaryServerInterceptor {
	//获取全局的tracer
	tracer := opentracing.GlobalTracer()

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var parentContext context.Context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		spanContext, err := tracer.Extract(opentracing.TextMap, MDReaderWriter{md})
		if err != nil {
			grpclog.Errorf("extract from metadata err: %v", err)
		} else {
			span := tracer.StartSpan(info.FullMethod, //name
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
				//标示类型的为grpc的Server
				ext.SpanKindRPCServer,
			)
			defer span.Finish()

			parentContext = opentracing.ContextWithSpan(ctx, span)
		}

		return handler(parentContext, req)
	}
}
