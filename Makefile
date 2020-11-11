SRC_ROOT=$(GOPATH)/src

deps:
	go get -t ./...

update-deps:
	go get -u ./...

protoc:
	protoc --proto_path=idl --go_out=$(SRC_ROOT) --go-grpc_out=$(SRC_ROOT) agency.proto agent.proto protocol.proto

install:
	@echo "Not implemented"

check_fmt:
	$(eval GOFILES = $(shell find . -name '*.go'))
	@gofmt -l $(GOFILES)

lint:
	@golangci-lint run

