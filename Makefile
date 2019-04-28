.DEFAULT_GOAL := help
.PHONY : check lint install-linters dep test 

OPTS=GO111MODULE=on 

lint: ## Run linters. Use make install-linters first.
	GO111MODULE=off vendorcheck ./...
	${OPTS} golangci-lint run ./...
	# The govet version in golangci-lint is out of date and has spurious warnings, run it separately
	${OPTS} go vet -all ./...

install: dep
	${OPTS} go build -o ~/Desktop/github.com/atang152/learn_listener_and_conn/basic_tcp_server ./cmd/basic_tcp_server/

install-linters: ## Install linters
	${OPTS} go get -u github.com/FiloSottile/vendorcheck
	# For some reason this install method is not recommended, see https://github.com/golangci/golangci-lint#install
	# However, they suggest `curl ... | bash` which we should not do
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	# GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.16.0
	${OPTS} go get -u golang.org/x/tools/cmd/goimports 

format: ## Formats the code. Must have goimports installed (use make install-linters).
	${OPTS} goimports -w -local github.com/atang152/learn_listener_and_conn ./pkg
	${OPTS} goimports -w -local github.com/atang152/learn_listener_and_conn ./cmd

dep: ## sorts dependencies
	${OPTS} go mod vendor -v

test: ## Run tests
	@mkdir -p coverage/
	${OPTS} go test -race -tags no_ci -cover -timeout=5m ./pkg/...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'