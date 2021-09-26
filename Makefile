
export PATH += $(`pwd`)/bin

GOBIN=$(`pwd`)/bin

PROTO_PATH=./proto
PROTO_OUTPUT=./proto

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

.PHONY: devs
devs:
	go install -tags=tools ./tools/tools.go
