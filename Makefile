SRC_ROOT=$(GOPATH)/src
IDL_PATH=./idl/v1

OUT_PATH ?= .grpc-out
DART_OUT_PATH ?= .dart-out

protoc_all: init protoc protoc_dart_all

protoc:	protoc_protocol protoc_agency protoc_agent protoc_authn

protoc_protocol:
	protoc --proto_path=$(IDL_PATH) --go_out=$(OUT_PATH) --go-grpc_out=$(OUT_PATH) protocol.proto

protoc_agency:
	protoc --proto_path=$(IDL_PATH) --go_out=$(OUT_PATH) --go-grpc_out=$(OUT_PATH) agency.proto

protoc_agent:
	protoc --proto_path=$(IDL_PATH) --go_out=$(OUT_PATH) --go-grpc_out=$(OUT_PATH) agent.proto

protoc_authn:
	protoc --proto_path=$(IDL_PATH) --go_out=$(OUT_PATH) --go-grpc_out=$(OUT_PATH) authn.proto

protoc_dart_protocol:
	protoc --proto_path=$(IDL_PATH) --dart_out=grpc:$(DART_OUT_PATH) -I$(IDL_PATH) protocol.proto

protoc_dart_agency:
	protoc --proto_path=$(IDL_PATH) --dart_out=grpc:$(DART_OUT_PATH) -I$(IDL_PATH) agency.proto

protoc_dart_agent:
	protoc --proto_path=$(IDL_PATH) --dart_out=grpc:$(DART_OUT_PATH) -I$(IDL_PATH) agent.proto

protoc_dart_authn:
	protoc --proto_path=$(IDL_PATH) --dart_out=grpc:$(DART_OUT_PATH) -I$(IDL_PATH) authn.proto

protoc_dart_all: protoc_dart_protocol protoc_dart_agency protoc_dart_agent protoc_dart_authn

clean:
	@rm -r $(OUT_PATH)/*
	@rm -r $(DART_OUT_PATH)/*

init:
	@echo "initialize build dirs:"
	@mkdir -pv $(OUT_PATH)
	@mkdir -pv $(DART_OUT_PATH)

check_fmt:
	$(eval GOFILES = $(shell find . -name '*.go'))
	@gofmt -l $(GOFILES)

lint:
	@golangci-lint run

