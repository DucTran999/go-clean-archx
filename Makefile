.PHONY: default help lint up down run deps
default: help

help: ## Show help for each of the Makefile commands
	@awk 'BEGIN \
		{FS = ":.*##"; printf "Usage: make ${cyan}<command>\n${white}Commands:\n"} \
		/^[a-zA-Z_-]+:.*?##/ \
		{ printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } \
		/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' \
		$(MAKEFILE_LIST)

lint: ## Run linters
	golangci-lint run --timeout 10m --config .golangci.yml

up: ## Setup testenv and migrate db
	@if [ ! -f .env ]; then \
		echo "❌ File .env not found. Please create one first."; \
		exit 1; \
	fi
	docker-compose up -d

down: ## Setup testenv and migrate db
	@if [ ! -f .env ]; then \
		echo "❌ File .env not found. Please create one first."; \
		exit 1; \
	fi
	docker-compose down

run: ## start the app locally
	go run cmd/main.go

deps: ## install library for generating mocks and merge code coverage
	go install github.com/vektra/mockery/v3@v3.4.0
	go install github.com/wadey/gocovmerge@latest

unit_test: ## Run all unit tests with coverage and save report to test/coverage/
	mkdir -p test/coverage
	go test -v -coverprofile=test/coverage/coverage.out ./internal/...

mock: ## generate mock
	mockery

codecov: ## check code coverage
	go tool cover -func=test/coverage/coverage.out
	go tool cover -html=test/coverage/coverage.out -o test/coverage/coverage.html
	
