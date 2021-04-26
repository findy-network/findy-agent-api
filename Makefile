SRC_ROOT=$(GOPATH)/src
IDL_PATH=idl/v1
OUT_PATH=.grpc-out/

deps:
	go get -t ./...

update-deps:
	go get -u ./...

protoc_protocol:
	protoc --proto_path=$(IDL_PATH) --go_out=$(SRC_ROOT) --go-grpc_out=$(OUT_PATH) protocol.proto

protoc_agency:
	protoc --proto_path=$(IDL_PATH) --go_out=$(SRC_ROOT) --go-grpc_out=$(OUT_PATH) agency.proto

protoc_agent:
	protoc --proto_path=$(IDL_PATH) --go_out=$(SRC_ROOT) --go-grpc_out=$(OUT_PATH) agent.proto

protoc:	protoc_protocol protoc_agency protoc_agent

clean:
	@rm -r $(OUT_PATH)*

install:
	@echo "Not implemented"

check_fmt:
	$(eval GOFILES = $(shell find . -name '*.go'))
	@gofmt -l $(GOFILES)

lint:
	@golangci-lint run

