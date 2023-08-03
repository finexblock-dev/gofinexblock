# gofinexblock

## Overview

Go 기반 application을 위한 모놀리식 레포지토리입니다. 

`cmd` 디렉토리에는 각각의 application들이 위치해있습니다. 

> `cmd`/```${application_name}```/`internal` 에는 application 내부에서만 사용될 수 있는 패키지들이 위치해있습니다.

`pkg` 디렉토리에는 공용으로 사용하는 패키지들이 위치해있습니다.


## Build

```shell
## build all applications
make build

## build specific application
make ${application}

## show help
make help
```

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

## Packages

패키지가 추가되면 README.md에 해당 패키지에 대한 설명을 추가해주세요.

### [entity](pkg/entity/entity.md)

### [admin](pkg/admin/admin.md)

### [announcement](pkg/announcement/announcement.md)

### [auth](pkg/auth/auth.md)

### [btcd](pkg/btcd/btcd.md)

### [cache](pkg/cache/cache.md)

### [compiler](pkg/compiler/compiler.md)

### [constant](pkg/constant/constant.md)

### [contracts](pkg/contracts/contracts.md)

### [daemon](pkg/daemon/daemon.md)

### [database](pkg/database/database.md)

### [engine](pkg/engine/engine.md)

### [entity](pkg/entity/entity.md)

### [ethereum](pkg/ethereum/ethereum.md)

### [files](pkg/files/files.md)

### [gen](pkg/gen/gen.md)

### [goaws](pkg/goaws/goaws.md)

### [goredis](pkg/goredis/goredis.md)

### [image](pkg/image/image.md)

### [instance](pkg/instance/instance.md)

### [interceptor](pkg/interceptor/interceptor.md)

### [order](pkg/order/order.md)

### [orderbook](pkg/orderbook/orderbook.md)

### [proto](pkg/proto/proto.md)

### [safety](pkg/safety/safety.md)

### [secure](pkg/secure/secure.md)

### [stream](pkg/stream/stream.md)

### [trade](pkg/trade/trade.md)

### [types](pkg/types/types.md)

### [user](pkg/user/user.md)

### [utils](pkg/utils/utils.md)

### [wallet](pkg/wallet/wallet.md)