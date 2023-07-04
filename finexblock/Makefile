.PHONY: build run-proxy run-ethereum run-polygon run-bitcoin bitcoin polygon ethereum proxy clean all proto vendor gen
GOCMD=go

GEN = gen
PROTO_PATH = proto

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)


## Vendor:
vendor: ## Run `go mod vendor`
	@go mod vendor

## Proto:
proto: ## Compile protobuf files
	@echo "${YELLOW}Compiling protobuf...${YELLOW}"
	@mkdir -p $(GEN)
	@protoc --go_out=$(GEN) --go-grpc_out=$(GEN) ./$(PROTO_PATH)/*/*.proto
	@echo "${CYAN}Compile done${RESET}"

## Clean:
clean: ## Clear all generated files
	@echo "${YELLOW}Clean generated files...${YELLOW}"
	@rm -rf $(GEN)/*
	@echo "${CYAN}Clean done${CYAN}"

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
			if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
			else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
			}' $(MAKEFILE_LIST)