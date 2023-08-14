# gofinexblock

## Overview

Go 기반 application을 위한 모놀리식 레포지토리입니다. 

`cmd` 디렉토리에는 각각의 application들이 위치해있습니다. 

> `cmd`/```${application_name}```/`internal` 에는 application 내부에서만 사용될 수 있는 패키지들이 위치해있습니다.

`pkg` 디렉토리에는 공용으로 사용하는 패키지들이 위치해있습니다.

## Build

### Make

모든 application은 `Makefile`을 통해서 컴파일됩니다. 

아래는 명령어 예시입니다. 

```shell
## build all applications
make build

## build specific application
make ${application}

## show help
make help
```

### Docker

Dockerfile들은 전부 `build` 디렉토리에 위치합니다. 

> 2023-08-03 현재 EC2에서 ECS로 이전작업 중입니다.  
> Dockerfile로 빌드한 이미지를 ECR에 push 후 ECS에서 pull하여 사용해야 할 듯 합니다.  
> 현재 EC2에서 사용하던 buildspec, appspec은 `scripts` 디렉토리에 위치하고 있습니다.   
> ECS에서 사용할 buildspec 및 기타 파일들이 있다면 `deployments` 디렉토리에 작업 부탁드립니다.

## Deploy

배포에 필요한 `buildspec`, `appspec`은 [deployments](deployments) 디렉토리에 위치합니다.

## Branch

작업 있을 시 Github issue 제작 후 `Create branch` 이용해서 이슈 번호로 분기해주세요.

### 반드시 dev 브랜치, 혹은 그 하위 브랜치에서 분기해주세요.

- master: 운영서버 배포용 브랜치
    - release->master 로 merge후 master 브랜치에서 각 application의 배포용 브랜치로 merge
- release: 개발서버 배포용 브랜치
- ${application_name}/release: 각 application의 개발서버 배포용 브랜치
    - dev->release 로 merge후 release 브랜치에서 각 application의 배포용 브랜치로 merge
- dev: 개발용 브랜치
  - feat/fix/refac/... -> dev로 merge


### `dev`, `release`, `master` 브랜치는 절대 빌드가 깨져서는 안됩니다.

> 2021-08-03 현재 빌드 자동화 작업 중입니다.  
> TODO: husky hook을 통해서 테스트와 빌드를 자동화할 예정입니다.


## Applications

### Backoffice

어드민 페이지 API 서버입니다. 

### Bitcoin daemon

비트코인 입출금 데몬 

### Bitcoin key

비트코인 signing 서버 

### Ethereum daemon

이더리움 입출금 데몬

### Ethereum key

이더리움 signing 서버

### Polygon daemon

폴리곤 입출금 데몬 

### Polygon key

폴리곤 signing 서버

### Proxy

Wallet proxy server, signing 서버로의 직접 호출을 막고, 모든 호출을 인터셉트하여 필요한 정보를 추가 혹은 로그를 기록한 뒤 전달합니다.

### Event subscriber

체결엔진으로부터 거래 이벤트를 hook으로 수신하고 데이터베이스에 반영합니다.

### Matching engine

trading server로 부터 지정가/시장가 주문 등록/취소 요청을 받고, 주문을 체결합니다.

오더북을 관리하고, 오더북 조회 및 자체 스냅샷 기능을 제공합니다.

## Refactoring

### 2023-08-11

#### TODO

1. 공통으로 사용하는 패키지에서 비즈니스 로직을 제거하고 각 application에서 구현하도록 변경.
2. 공통으로 사용하는 패키지에서 재사용이 되는 로직들은 합쳐서 하나의 패키지로 사용하도록 변경.
3. fiber => gin으로 변경(net/http 패키지와 호환)