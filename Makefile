include .env
LOCAL_BIN:=$(CURDIR)/bin

lint:
	golangci-lint run ./... --config .golangci.yaml
PHONY: lint

lint-fix:
	golangci-lint run ./... --config .golangci.yaml --fix
PHONY: lint-fix

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
PHONY: install-deps

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
PHONY: get-deps

generate: generate-user-api  generate-auth-api generate-access-api
PHONY: generate

generate-user-api:
	mkdir -p pkg/userApi
	mkdir -p pkg/swagger
	protoc --proto_path api/userApi --proto_path vendor.protogen \
	--go_out=pkg/userApi --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/userApi --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--grpc-gateway_out=pkg/userApi --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/userApi/userApi.proto

PHONY: generate-user-api

generate-auth-api:
	mkdir -p pkg/auth
	protoc --proto_path api/auth \
	--go_out=pkg/auth --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/auth --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/auth/auth.proto

PHONY: generate-auth-api

generate-access-api:
	mkdir -p pkg/access
	protoc --proto_path api/access \
	--go_out=pkg/access --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/access --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/access/access.proto

PHONY: generate-access-api

vendor-proto:
		if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi
PHONY: vendor-proto

migrate-down:
	migrate -path=migrations -database=postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable down
PHONY: migrate-down

migrate-up:
	migrate -path=migrations -database=postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable up
PHONY: migrate-up

# запускать в корне проекта (run in the root of the project)
run: format
	go run cmd/main.go
PHONY: run

start-app: format
	docker-compose up --build -d
PHONY: start-app

stop-app:
	docker-compose down
PHONY: stop-app

format:
	go fmt ./...
PHONY: format