PROJECT_NAME := "openvagas-api"


.PHONY: all dep build clean lint


setup: ## Setup project
	@go install golang.org/x/vuln/cmd/govulncheck@latest

all: build

dep-dev-run: ## Run development dependencies
	@docker-compose -f docker-compose.yml up -d  --build --remove-orphans

dep-dev-stop: ## Stop development dependencies
	@docker-compose stop

dep-dev-status: ## Show status from development dependencies
	@docker-compose status

dev-start-with-tools: dep-dev-run ## Run application and dependencies
	@docker-compose -f docker-compose.app.yml -f docker-compose.tools.yml -f docker-compose.yml up -d  --build --remove-orphans

dev-start-with-db: dep-dev-run ## Run application and dependencies
	@docker-compose -f  docker-compose.yml up -d  --build --remove-orphans

dev-start-db: dep-dev-run ## Run application and dependencies
	@docker-compose up -d  --build --remove-orphans

show-logs: ## Show logs from development dependencies
	@docker logs ngps-api -f --tail=100

create-network-external:
	@docker network create openvagas

create-network:
	@docker network create openvagas

lint: ## Lint the files
	@staticcheck ./...

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file
	@go build -v -o bin/${PROJECT_NAME} web/main.go

clean: ## Remove previous build
	@rm -f bin/$(PROJECT_NAME)

test-unit: ## run tests
	mkdir -p coverage
	go test --tags=unit -v -coverprofile ./coverage/cover.out ./...

test-e2e: ## run e2e tests
	go test --tags=e2e -v ./...

test-e2e-clear-cache: ## clear test cache 
	go clean -testcache

test-e2e-local: test-e2e-clear-cache ## run e2e tests local
	APPLICATION_URL=http://localhost:8080 \
	APPLICATION_VERSION= \
	TEST_TIMEOUT=0 \
	go test --tags=e2e -v ./...

showcover: unit ## show tests results
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
	open ./coverage/cover.html

help: ## Show commands and description
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

security: ## Check security vulnerabilities
	@govulncheck ./...