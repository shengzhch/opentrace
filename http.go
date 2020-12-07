package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc/metadata"
)

func Normal() {
	g := opentracing.GlobalTracer()
	span := g.StartSpan("normal_start")
	defer span.Finish()
	time.Sleep(time.Second)

	//把span放到ctx中
	spanCtx := opentracing.ContextWithSpan(context.Background(), span)

	//ChildOf
	s1, ctx1 := opentracing.StartSpanFromContext(spanCtx, "normal_start_step1")

	//等价于以下写法
	//s1 := opentracing.StartSpan("normal_start_step1", opentracing.ChildOF(span.Context()))
	//ctx1 := opentracing.ContextWithSpan(spanCtx, s1)

	defer s1.Finish()
	time.Sleep(time.Second)

	s2, _ := opentracing.StartSpanFromContext(ctx1, "normal_start_step2")
	defer s2.Finish()
	time.Sleep(time.Second)
}

func Follow() {
	g := opentracing.GlobalTracer()
	span := g.StartSpan("follow_start")
	time.Sleep(time.Second)
	span.Finish()
	s1 := opentracing.StartSpan("follow_start_step1", opentracing.FollowsFrom(span.Context()))
	s1.Finish()
	time.Sleep(time.Second)
}

func Child() {
	g := opentracing.GlobalTracer()
	span := g.StartSpan("child_start")
	defer span.Finish()
	time.Sleep(time.Second)
	s1 := opentracing.StartSpan("child_start_step1", opentracing.ChildOf(span.Context()))
	defer s1.Finish()
	time.Sleep(time.Second)
}

func Rpc() {
	g := opentracing.GlobalTracer()
	spanhttp := g.StartSpan("http_rpc")
	defer spanhttp.Finish()

	spanHttpCtx := opentracing.ContextWithSpan(context.Background(), spanhttp)
	callRpcSpan, ctx := formatRpcCtxWithSpan(spanHttpCtx)
	_, err := rpcCli.Welcome(ctx, &WelcomeReq{})
	callRpcSpan.Finish()

	if err != nil {
		fmt.Println("rpcCli Welcome Failed ", err)
		return
	}
}

func formatRpcCtxWithSpan(c context.Context) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(c, "call_rpc",
		opentracing.Tag{Key: string(ext.Component), Value: "gRPC"},
		ext.SpanKindRPCClient)

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	} else {
		md = md.Copy()
	}
	err := opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, MDReaderWriter{md})
	if err != nil {
		span.LogFields(log.String("inject-error", err.Error()))
	}

	return span, metadata.NewOutgoingContext(ctx, md)
}

func Do(f func()) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f()
		_, _ = w.Write([]byte(r.URL.Path))
	}
}

func httpRun() {
	http.HandleFunc("/normal", Do(Normal))
	http.HandleFunc("/flow", Do(Follow))
	http.HandleFunc("/child", Do(Child))

	http.HandleFunc("/rpc", Do(Rpc))
	_ = http.ListenAndServe(":1200", nil)
	fmt.Println("server done ")
}
