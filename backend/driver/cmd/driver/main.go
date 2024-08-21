package main

import (
	"flag"
	"os"

	"driver/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	capi "github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "DriverService"
	// Version is the version of the compiled software.
	Version string = "0.0.1"
	// flagconf is the config flag.
	flagconf string = "../../configs"

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(cs *conf.Service, logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	reg := initServiceRegistry(cs.Consul.Address)
	err := initTrace(cs.Jaeger.Address)
	if err != nil {
		log.Fatal(err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(reg),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Service, bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func initServiceRegistry(address string) *consul.Registry {

	// 1，获取consul客户端
	consulConfig := capi.DefaultConfig()

	// 通过配置文件拿到consul服务的地址
	consulConfig.Address = address
	consulClient, err := capi.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 2，创建consul注册器
	consulRegister := consul.New(consulClient)

	return consulRegister
}

// 初始化Tracer
// @param url string 指定 Jaeger 的API接口
// ：14268/api/traces
func initTrace(url string) error {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}

	traceProvider := traceSDK.NewTracerProvider(
		// 采样设置
		traceSDK.WithSampler(traceSDK.AlwaysSample()),
		// 使用 exporter 作为批处理程序
		traceSDK.WithBatcher(exporter),
		// 将当前服务信息，作为资源告知给TracerProvider
		traceSDK.WithResource(resource.NewSchemaless(
			// 必要的配置
			semconv.ServiceNameKey.String("Valuation"),
			// 任意的其他属性配置
			attribute.String("exporter", "jaeger"),
		)),
	)
	otel.SetTracerProvider(traceProvider)
	return nil
}
