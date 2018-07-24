PROTO_PATH=proto

all: protobuf build

build:
	go build

clean:
	rm -vf $(PROTO_PATH)/*.go

protobuf:
	protoc -I $(PROTO_PATH)/ --go_out=plugins=grpc:$(PROTO_PATH)/ $(PROTO_PATH)/*.proto
