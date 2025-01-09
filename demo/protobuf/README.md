## Protocol Buffers

Language Guide (proto3): https://developers.google.com/protocol-buffers/docs/proto3

Protocol Buffer Basics: Go: https://developers.google.com/protocol-buffers/docs/gotutorial

## 安装 protoc

https://github.com/protocolbuffers/protobuf/releases

### MAC

下载 protoc-3.13.0-osx-x86_64.zip

```bash
unzip protoc-3.13.0-osx-x86_64
sudo mv protoc /usr/local
```

~/.zshrc 文件添加

```text
export PATH="/usr/local/protoc/bin:$PATH"
```

## 安装 protoc-gen-go

$GOPATH 一般为 /Users/x/go

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
cd {$GOPATH}/pkg/mod/github.com/golang/protobuf@v1.4.2/protoc-gen-go
sudo go install
ls {$GOPATH}/bin
```

