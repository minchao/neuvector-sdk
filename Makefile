SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: download-api-spec
download-api-spec: ## Download OpenAPI spec
	@curl -s https://raw.githubusercontent.com/neuvector/neuvector/main/controller/api/apis.yaml -o apis.yaml

.PHONY: validate-api-spec
validate-api-spec: install.swagger ## Validate OpenAPI spec
	@swagger version
	swagger validate ./apis.yaml

.PHONY: lint
lint: install.golangci-lint ## Run linter
	@golangci-lint version
	golangci-lint run ./...

fix-api-spec:
	$(MAKE) fix-field-and-method-with-the-same-name-validate

fix-field-and-method-with-the-same-name-validate:
	yq eval -i '.definitions.RESTAdmissionState.properties.ctrl_states.properties.validate."x-go-name" = "ValidateType"' apis.yaml

##@ Build

.PHONY: generate
generate: install.swagger validate-api-spec fix-api-spec ## Generate code
	@swagger version
	swagger generate client -f ./apis.yaml

# Tools

.PHONY: install.swagger
install.swagger:
ifeq (, $(shell which swagger))
	$(error The 'swagger' command not found, install it from https://goswagger.io/install.html)
endif

.PHONY: install.golangci-lint
install.golangci-lint:
ifeq (, $(shell which golangci-lint))
	$(error The 'golangci-lint' command not found, install it from https://golangci-lint.run/usage/install/)
endif

.PHONY: install.yq
install.yq:
ifeq (, $(shell which yq))
	$(error The 'yq' command not found, install it from https://formulae.brew.sh/formula/yq)
endif
