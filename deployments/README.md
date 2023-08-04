# Deployments

## Overview

`appspec.yml`, `buildspec.yml` 등 배포에 필요한 파일들이 위치합니다.

경로는 다음과 같이 설정 부탁드립니다. 

```shell
touch deployments/$(application)/buildspec.yml
```

## ECS Cluster

### finexblock-wallet-server

####  domain: https://bitcoin.finexblock.com
####  domain: https://polygon.finexblock.com
####  domain: https://ethereum.finexblock.com
####  domain: https://proxy.finexblock.com

####  domain: https://bitcoin-dev.finexblock.com 
####  domain: https://ethereum-dev.finexblock.com
####  domain: https://polygon-dev.finexblock.com
####  domain: https://proxy-dev.finexblock.com

#### port: 50051

Services 

- wallet-proxy-stage/prod
- bitcoin-key-stage/prod
- ethereum-key-stage/prod
- polygon-key-stage/prod

### finexblock-matching-engine

####  domain: ${symbol}-matching-engine-dev.finexblock.com
####  domain: ${symbol}-matching-engine.finexblock.com

#### port: 50051

- btc-eth-matching-engine-stage/prod
- btc-matic-matching-engine-stage/prod
- btc-etc-matching-engine-stage/prod
- btc-sand-matching-engine-stage/prod

### finexblock-core-server

####  domain: https://core-dev.finexblock.com
####  domain: https://core.finexblock.com

#### port: 80, 443

#### port: 80, 443

- core-server-stage/prod

### finexblock-trading-server

####  domain: https://trading-dev.finexblock.com

#### port: 80, 443, 50051 

- trading-server-stage/prod

### finexblock-backoffice

####  domain: https://bacoffice-api-dev.finexblock.com
####  domain: https://bacoffice-api.finexblock.com

#### port: 80, 443 

- backoffice-stage/prod


## EC2 instance 

### finexblock-event-subscriber(2)

### btc-node(2)

### btc-key-node(2)

### finexblock-wallet-daemon(6)

### finexblock-bastion-host(2)

### finexblock-mysql(2)