# protobufModels
Storage of protobuf .proto files that generate *.go files independent of the implementation projects. Helps keep the requirement out of the other projects build.

# prerequisites
to assure that the tools are already installed for protobufs, i recommend 
- using go 1.12+
- unset GOPATH from environment variables
- go get -u google.golang.org/grpc
- go get -u github.com/golang/protobuf/protoc-gen-go
- add the newly compiled go plugin to your path
  export PATH=$PATH:~/go/bin //default install on OSX

# modify .proto files
anytime you modify the protobuf definition files, just regen all the native go files 

```protoc -I . helloProto.proto --go_out=plugins=grpc:.```