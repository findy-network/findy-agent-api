SRC_ROOT=$(GOPATH)/src
IDL_PATH=idl/v1

deps:
	go get -t ./...

update-deps:
	go get -u ./...

protoc_protocol:
	protoc --proto_path=$(IDL_PATH) --go_out=$(SRC_ROOT) --go-grpc_out=$(SRC_ROOT) protocol.proto

protoc_agency:
	protoc --proto_path=$(IDL_PATH) --go_out=$(SRC_ROOT) --go-grpc_out=$(SRC_ROOT) agency.proto

protoc_agent:
	protoc --proto_path=$(IDL_PATH) --go_out=$(SRC_ROOT) --go-grpc_out=$(SRC_ROOT) agent.proto

protoc:	protoc_protocol protoc_agency protoc_agent

install:
	@echo "Not implemented"

check_fmt:
	$(eval GOFILES = $(shell find . -name '*.go'))
	@gofmt -l $(GOFILES)

lint:
	@golangci-lint run

