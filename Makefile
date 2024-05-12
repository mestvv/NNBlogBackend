LOCAL_BIN:=$(CURDIR)/bin
BUILD_DIR:=$(CURDIR)/cmd/app

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.yaml

# compose deps
compose:
	@echo 'compose deps'
	docker compose -f dev/docker-compose.yaml up -d

# down deps
compose-down:
	@echo 'compose deps'
	docker compose -f dev/docker-compose.yaml down 

build: deps build_binary

build-binary:
	@echo 'build backend binary'
	go build -o $(LOCAL_BIN) $(BUILD_DIR)

deps:
	@echo 'install dependencies'
	go mod tidy -v

run: deps run-app

run-app:
	@echo 'run backend'
	go run $(BUILD_DIR)/main.go

# generate swagger
swag:
	@echo 'generation swagger docs'
	swag init --parseDependency -g handler.go -dir internal/api/internal/http/v1 --instanceName internal
