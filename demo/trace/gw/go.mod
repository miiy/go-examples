module github.com/miiy/go-examples/trace/gw

go 1.18

replace (
	github.com/miiy/go-examples/trace/api => ../api
	github.com/miiy/go-examples/trace/common => ../common
)

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/miiy/go-examples/trace/api v0.0.0-00010101000000-000000000000
	github.com/miiy/go-examples/trace/common v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.35.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.35.0
	go.opentelemetry.io/contrib/propagators/b3 v1.10.0
	go.opentelemetry.io/otel v1.10.0
	google.golang.org/grpc v1.49.0
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.10.0 // indirect
	go.opentelemetry.io/otel/sdk v1.10.0 // indirect
	go.opentelemetry.io/otel/trace v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
	golang.org/x/sys v0.0.0-20220913175220-63ea55921009 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
