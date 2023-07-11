.PHONY: build run-proxy run-ethereum run-polygon run-bitcoin bitcoin polygon ethereum proxy clean all proto vendor gen
GOCMD=go

GEN = gen
PROTO_PATH = proto
CONTRACT_PATH = contracts

## Contract abis
ABI = $(GEN)/abi
BIN = $(GEN)/bin

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
	@protoc --go_out=$(GEN) --go-grpc_out=$(GEN) ./$(PROTO_PATH)/**/*.proto
	@echo "${CYAN}Compile done${RESET}"

## Abi:
abi:
	@echo "${YELLOW}Generating ABIs...${YELLOW}"
	@mkdir -p $(ABI)
	@solc --abi --overwrite -o $(ABI) $(CONTRACT_PATH)/*.sol
	@echo "${CYAN}Generate done${CYAN}"

bin:
	@echo "${YELLOW}Generating Binary files...${YELLOW}"
	@mkdir -p $(BIN)
	@solc --bin --overwrite -o $(BIN) $(CONTRACT_PATH)/*.sol
	@echo "${CYAN}Generate done${CYAN}"

## Abigen:
abigen: abi bin ## Compile solidity files and generate abi
	@echo "${YELLOW}Compiling smart contract...${YELLOW}"
	@mkdir -p $(GEN)/$(CONTRACT_PATH)
	@abigen --abi $(ABI)/*.abi --bin $(BIN)/*.bin --pkg=ethereum --out=$(GEN)/$(CONTRACT_PATH)/gofinexblock.go
	@echo "${CYAN}Compile done${RESET}"

## Contract:
contract: ## Compile solidity files
	@echo "${YELLOW}Compiling smart contract...${YELLOW}"
	@mkdir -p $(GEN)
	@solc --abi --bin --overwrite --pkg=gofinexblock $(GEN) $(CONTRACT_PATH)/**/*.sol
	@echo "${CYAN}Compile done${RESET}"

## Clean:
clean: ## Clear all generated files
	@echo "${YELLOW}Clean generated files...${YELLOW}"
	@rm -rf $(GEN)/*
	@echo "${CYAN}Clean done${CYAN}"

## Build:
build: proto

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