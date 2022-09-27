.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this message.
## -----------------------------------------------------------------------------
		@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

.PHONY: ci
## ci:          Run ci.
ci:
		make lint
		make test

.PHONY: test
## test:        Run all tests in project.
test:
		go test -v -race -cover -bench=. ./...

.PHONY: lint
## lint:        Run golangci-lint for project.
lint:
		golangci-lint run -v

.PHONY: lint_docker
## lint_docker: Run golangci-lint for project.
lint_docker:
		docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.49.0 golangci-lint run -v
