export GOBIN=$(shell pwd)/bin
export PATH += $(GOBIN)

PROTO_PATH=./proto
PROTO_OUTPUT=./proto

.DEFAULT_GOAL=build

.PHONY: protoc_gen
proto_gen:
	@protoc --proto_path=$(PROTO_PATH) \
		--go_out=$(PROTO_OUTPUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUTPUT) --go-grpc_opt=paths=source_relative \
		--plugin=./bin/protoc-gen-go --plugin=./bin/protoc-gen-go-grpc \
		$(PROTO_PATH)/A/api.proto
	@echo "Generated source files from $(PROTO_PATH)/A/api.proto."
	@protoc --proto_path=$(PROTO_PATH) \
		--go_out=$(PROTO_OUTPUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUTPUT) --go-grpc_opt=paths=source_relative \
		--plugin=./bin/protoc-gen-go --plugin=./bin/protoc-gen-go-grpc \
		$(PROTO_PATH)/B/api.proto
	@echo "Generated source files from $(PROTO_PATH)/B/api.proto."
	@protoc --proto_path=$(PROTO_PATH) \
		--go_out=$(PROTO_OUTPUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUTPUT) --go-grpc_opt=paths=source_relative \
		--plugin=./bin/protoc-gen-go --plugin=./bin/protoc-gen-go-grpc \
		--descriptor_set_out=$(PROTO_OUTPUT)/proxy/api.protoset --include_imports \
		$(PROTO_PATH)/proxy/api.proto
	@echo "Generated source files from $(PROTO_PATH)/proxy/api.proto."

clean:
	find ./proto -type f -name '*.go' | /usr/bin/xargs rm
	rm -rf ./bin/backend_a ./bin/backend_b ./bin/proxy

.PHONY: build_proxy
build_proxy: proto_gen
	go build -o ./bin/proxy ./cmd/proxy/main.go

.PHONY: build_a
build_a: proto_gen
	go build -o ./bin/backend_a ./cmd/A/main.go

.PHONY: build_b
build_b: proto_gen
	go build -o ./bin/backend_b ./cmd/B/main.go

.PHONY: build
build: build_a build_b build_proxy

.PHONY: dev-deps
dev-deps: deps
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: deps
deps:
	go mod tidy
	go mod vendor

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yml ./...
