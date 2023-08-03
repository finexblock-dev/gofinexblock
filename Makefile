.PHONY: build run-proxy run-ethereum run-polygon run-bitcoin \
bitcoin-key polygon-key ethereum-key proxy \
backoffice run-backoffice \
run-bitcoin run-ethereum run-polygon run-proxy \
run-bitcoin-daemon run-ethereum-daemon run-polygon-daemon \
run-event-subscriber run-matching-engine \
proto abi bin erc20 \
clean all vendor

GOCMD = $$(which go)

INIT = init
CMD = cmd
PKG = pkg
GEN = gen
CONTRACT = contracts
PROTO = proto


ABI = $(GEN)/abi
BIN = $(GEN)/bin


ETHEREUM_KEY = ethereum_key
POLYGON_KEY = polygon_key
BITCOIN_KEY = bitcoin_key
PROXY = proxy

ETHEREUM_DAEMON = ethereum_daemon
POLYGON_DAEMON = polygon_daemon
BITCOIN_DAEMON = bitcoin_daemon

MATCHING_ENGINE = matching_engine
EVENT_SUBSCRIBER = event_subscriber
BACKOFFICE = backoffice

ETHEREUM_KEY_OUTFILE = ethereum-key
POLYGON_KEY_OUTFILE = polygon-key
BITCOIN_KEY_OUTFILE = bitcoin-key
PROXY_OUTFILE = proxy

ETHEREUM_DAEMON_OUTFILE = ethereum-daemon
POLYGON_DAEMON_OUTFILE = polygon-daemon
BITCOIN_DAEMON_OUTFILE = bitcoin-daemon

MATCHING_ENGINE_OUTFILE = matching-engine
EVENT_SUBSCRIBER_OUTFILE = event-subscriber
BACKOFFICE_OUTFILE = backoffice

ENTRY = main.go

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)


all: help

## Build:
build: ethereum-key polygon-key bitcoin-key ethereum-daemon polygon-daemon bitcoin-daemon event-subscriber matching-engine proxy backoffice proto abi bin erc20 ## Build all servers

## Vet:
vet: ## Run `go vet cmd/main.go`
	@$(GOCMD) vet $(CMD)/$(ENTRY)

## Vendor:
vendor: ## Run `go mod vendor`
	@go mod vendor

## Proto:
proto: ## Compile protobuf files
	@echo "${YELLOW}Compiling protobuf...${YELLOW}"
	@mkdir -p $(PKG)/$(GEN)
	@protoc --go_out=$(PKG)/$(GEN) --go-grpc_out=$(PKG)/$(GEN) $(PKG)/$(PROTO)/**/*.proto
	@echo "${CYAN}Compile done${RESET}"

## Abi:
abi: ## Generate abi files
	@echo "${YELLOW}Generating ABIs...${YELLOW}"
	@mkdir -p $(PKG)/$(ABI)
	@solc --abi --overwrite -o $(PKG)/$(ABI) $(PKG)/$(CONTRACT)/*.sol
	@echo "${CYAN}Generate done${CYAN}"

## Bin:
bin: ## Generate binary files
	@echo "${YELLOW}Generating Binary files...${YELLOW}"
	@mkdir -p $(PKG)/$(BIN)
	@solc --bin --overwrite -o $(PKG)/$(BIN) $(PKG)/$(CONTRACT)/*.sol
	@echo "${CYAN}Generate done${CYAN}"

## ERC20:
erc20: abi bin ## Compile solidity files and generate abi
	@echo "${YELLOW}Compiling smart contract...${YELLOW}"
	@mkdir -p $(PKG)/$(GEN)/$(CONTRACT)
	@abigen --abi $(PKG)/$(ABI)/*.abi --bin $(PKG)/$(BIN)/*.bin --pkg=ERC20 --out=$(PKG)/$(GEN)/$(CONTRACT)/ERC20.go
	@echo "${CYAN}Compile done${RESET}"


## Ethereum-Key
ethereum-key: ## Build Ethereum Key Server
	@echo "${YELLOW}Building ethereum server ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(ETHEREUM_OUTFILE) $(CMD)/$(ETHEREUM_KEY)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Polygon-Key
polygon-key: ## Build Polygon Key Server
	@echo "${YELLOW}Building polygon server ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(POLYGON_OUTFILE) $(CMD)/$(POLYGON_KEY)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Bitcoin-Key
bitcoin-key:  ## Build Bitcoin Key Server
	@echo "${YELLOW}Building bitcoin server ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(BITCOIN_OUTFILE) $(CMD)/$(BITCOIN_KEY)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Proxy
proxy: ## Build Wallet Proxy Server
	@echo "${YELLOW}Building wallet proxy server ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(PROXY_OUTFILE) $(CMD)/$(PROXY)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Bitcoin daemon
bitcoin-daemon: ## Build bitcoin daemon
	@echo "${YELLOW}Building bitcoin daemon ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(BITCOIN_DAEMON_OUTFILE) $(CMD)/$(BITCOIN_DAEMON)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Ethereum daemon
ethereum-daemon: ## Build ethereum daemon
	@echo "${YELLOW}Building ethereum daemon ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(ETHEREUM_DAEMON_OUTFILE) $(CMD)/$(ETHEREUM_DAEMON)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Polygon daemon
polygon-daemon: ## Build polygon daemon
	@echo "${YELLOW}Building polygon daemon ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(POLYGON_DAEMON_OUTFILE) $(CMD)/$(POLYGON_DAEMON)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Event subscriber
event-subscriber: ## Build event subscriber
	@echo "${YELLOW}Building event subscriber ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(EVENT_SUBSCRIBER_OUTFILE) $(CMD)/$(EVENT_SUBSCRIBER)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"


## Matching engine
matching-engine: ## Build matching engine
	@echo "${YELLOW}Building matching engine ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(MATCHING_ENGINE_OUTFILE) $(CMD)/$(MATCHING_ENGINE)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Backoffice swagger
backoffice-swagger: ## Generate backoffice swagger
	@swag init -g $(CMD)/$(BACKOFFICE)/$(ENTRY) -q --output=./$(CMD)/$(BACKOFFICE)/docs --parseInternal --parseVendor --generatedTime --requiredByDefault
	@#swag fmt .

## Backoffice
backoffice: backoffice-swagger ## Build backoffice server
	@echo "${YELLOW}Building backoffice server ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(BACKOFFICE_OUTFILE) $(CMD)/$(BACKOFFICE)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Run Proxy
run-proxy: proxy ## Run Wallet Proxy Server
	@$(INIT)/$(PROXY_OUTFILE)

## Run Bitcoin
run-bitcoin: bitcoin-key ## Run Bitcoin Key Server
	@$(INIT)/$(BITCOIN_OUTFILE)

## Run Ethereum
run-ethereum: ethereum-key ## Run Ethereum Key Server
	@$(INIT)/$(ETHEREUM_OUTFILE)

## Rum Polygon
run-polygon: polygon-key ## Run Polygon Key Server
	@$(INIT)/$(POLYGON_OUTFILE)

## Run Bitcoin daemon
run-bitcoin-daemon: bitcoin-daemon ## Run Bitcoin daemon
	@$(INIT)/$(BITCOIN_DAEMON_OUTFILE)

## Run Ethereum daemon
run-ethereum-daemon: ethereum-daemon ## Run Ethereum daemon
	@$(INIT)/$(ETHEREUM_DAEMON_OUTFILE)

## Run Polygon daemon
run-polygon-daemon: polygon-daemon ## Run Polygon daemon
	@$(INIT)/$(POLYGON_DAEMON_OUTFILE)

## Run Event subscriber
run-event-subscriber: event-subscriber ## Run Event subscriber
	@$(INIT)/$(EVENT_SUBSCRIBER_OUTFILE)

## Run Matching engine
run-matching-engine: matching-engine ## Run Matching engine
	@$(INIT)/$(MATCHING_ENGINE_OUTFILE)

## Clean:
clean: ## Clear all generated files
	@echo "${YELLOW}Clean generated files...${YELLOW}"
	@rm -rf $(BUILD)/*
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