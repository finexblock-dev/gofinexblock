.PHONY: build run-proxy run-ethereum run-polygon run-bitcoin bitcoin-key polygon-key ethereum-key proxy clean all vendor

## Commands
GOCMD = $$(which go)

## Directories
INIT = init
CMD = cmd

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

## Outfiles
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

## main function
ENTRY = main.go

## Utils
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)


all: help

## Build:
build: ethereum-key polygon-key bitcoin-key ethereum-daemon polygon-daemon bitcoin-daemon ## event-subscriber matching-engine proxy ## Build all servers

## Vet:
vet: ## Run `go vet cmd/main.go`
	@$(GOCMD) vet $(CMD)/$(ENTRY)

## Vendor:
vendor: ## Run `go mod vendor`
	@go mod vendor

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
	@echo "${YELLOW}Building ethereum daemon server ...${YELLOW}"
	@GO111MODULE=on \
	CGO_ENABLED=1 \
	GOARCH=amd64 \
	$(GOCMD) build -mod vendor -o $(INIT)/$(ETHEREUM_DAEMON_OUTFILE) $(CMD)/$(ETHEREUM_DAEMON)/$(ENTRY)
	@echo "${CYAN}Build done${CYAN}"

## Polygon daemon
polygon-daemon: ## Build polygon daemon
	@echo "${YELLOW}Building polygon daemon server ...${YELLOW}"
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

## Backoffice
backoffice: ## Build backoffice server
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