Articles linke:
    1. https://ewanvalentine.io/microservices-in-golang-part-1/

gRPC: Mechanism for maintaining distributed application servers and clients.
Protocol Buffer (ProtBuf): Mechanism for converting structured data to seralize data.

Install gRPC/protbuf
    1. Check version (>= 1.6)
        go version

    2. Install grpc
        go get -u google.golang.org/grpc

    3. Install protocol buffers v3
        - Download suitable compiler file from: 
            https://github.com/protocolbuffers/protobuf/releases
        - Set path to file this dir in $PATH
        - Install protoc plugin for go
            go get -u github.com/golang/protobuf/protoc-gen-go
        - Install protoc3
            https://gist.github.com/sofyanhadia/37787e5ed098c97919b8c593f0ec44d8
        

