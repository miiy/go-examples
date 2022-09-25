module github.com/miiy/go-examples/trace/order

go 1.18

replace (
	github.com/miiy/go-examples/trace/api => ../api
	github.com/miiy/go-examples/trace/common => ../common
)

require (
	github.com/miiy/go-examples/trace/api v0.0.0-00010101000000-000000000000
	github.com/miiy/go-examples/trace/common v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.35.0
	go.opentelemetry.io/contrib/propagators/b3 v1.10.0
	go.opentelemetry.io/otel v1.10.0
	google.golang.org/grpc v1.49.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.10.0 // indirect
	go.opentelemetry.io/otel/sdk v1.10.0 // indirect
	go.opentelemetry.io/otel/trace v1.10.0 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
