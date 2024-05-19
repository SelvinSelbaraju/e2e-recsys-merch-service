# Requirements

In order to convert the `.proto` files to the `pb.go` files in this repo, the below dependencies are required:

1. Protoc
```bash
brew install protobuf
```
2. Go plugins for Protoc
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

To prevent VSCode from complaining that the import of `model_config.proto` does not work, update the VSCode settings with protoc options to the proto path.
```json
...other settings
    "protoc": {
        "options": [
            "-I internal/proto",
        ]
    }
```

Using these dependencies, we copy and pasted the `.proto` files from [here](https://github.com/triton-inference-server/common/tree/main/protobuf). These `.proto` files should not be changed (except for adding a go_package option like `option go_package = "./proto"`), because the inference server uses strict definitions of the package and variable names. We then compile them using `protoc`:
```bash
# Work dir is root of the repo
# -I says where to look for protos for importing
# We need to generate go output and go grpc output
# Go output for compiling message related code
# Go grpc output for compiling service and rpc related code
protoc -I internal/proto --go-grpc_out="./internal" --go_out="./internal" internal/proto/*.proto
