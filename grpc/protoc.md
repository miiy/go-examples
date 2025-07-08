# protoc

## Install protoc

https://github.com/protocolbuffers/protobuf

install plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

clone google api

```bash
cd third_party
git clone https://github.com/googleapis/googleapis.git
```

## gerenate gRPC code

```bash
protoc --proto_path=api/tag/v1 --proto_path=third_party/proto --go_out=plugins=grpc:api/tag/v1 --go_opt paths=source_relative tag_service.proto
    example proto
    @protoc -I ./api/example/proto/v1 -I ./third_party/proto \
            -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 \
            --go_out ./pkg/api/example/proto/v1 --go_opt paths=source_relative \
            --go-grpc_out ./pkg/api/example/proto/v1 --go-grpc_opt paths=source_relative \
            --grpc-gateway_out ./pkg/api/example/proto/v1 --grpc-gateway_opt paths=source_relative \
            --validate_out="lang=go:./pkg/api/example/proto/v1" \
            ./pkg/api/example/proto/v1/example_service.proto
```

## References

Protocol Buffers Documentation: <https://protobuf.dev/>

gRPC: <https://grpc.io/>

<https://grpc.io/docs/languages/go/quickstart/>

<https://github.com/protocolbuffers/protobuf-go>

<https://github.com/grpc/grpc-go>
