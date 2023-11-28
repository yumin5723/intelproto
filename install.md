# go-micro环境安装

## 环境要求
- Go 1.19+

## 环境安装步骤

1. 在本地开发环境安装protobuf，推荐版本v25.1。下载地址：[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)

2. 安装`protoc-gen-go` 和 `protoc-gen-micro` 插件。请严格按照以下命令安装，因为本项目使用的是go-micro v3版本，proto相关插件必须安装适配的版本。

```bash
go install github.com/asim/go-micro/cmd/protoc-gen-micro/v3@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
```

3. 使用以下命令生成proto文件：

```bash
protoc --proto_path=. --micro_out=. --go_out=. greeter.proto
```

请将`greeter.proto`替换为你自己的原始proto文件名。