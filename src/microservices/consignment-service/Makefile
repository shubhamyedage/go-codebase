build:
	protoc proto/consignment/consignment.proto --go_out=plugins=grpc:./
	# protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/microservices/consignment-service/proto/consignment/consignment.proto