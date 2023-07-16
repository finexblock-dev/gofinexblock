# orderbook

## Description

Orderbook을 관리하고 저장할 수 있게 해주는 모듈입니다.

[`Queue`](interface.go), [`Service`](interface.go), [`Repository`](interface.go)의 세 가지 인터페이스를 이용하여 구현되어 있습니다.

사용자는 `Queue`를 통해서 주문에 대한 요청을 할 수 있습니다. 

한 번 Queue에 접수된 요청은 Redis stream을 이용해서 처리됩니다. 

`Queue`는 내부 채널을 이용해서 주문을 접수하고, `Service`에게 주문에 대한 요청을 전달합니다.

## Stream

Redis stream을 이용해서 주문을 처리합니다. 

사용하는 Redis stream의 key는 다음과 같습니다. 

- `STREAM:MATCH`: 주문이 매칭되었을 때 발생하는 stream 입니다. 다음 12개의 경우로 나뉘어 처리됩니다.
  - `CASE:LIMIT_ASK_BIGGER`
  - `CASE:LIMIT_ASK_EQUAL` 
  - `CASE:LIMIT_ASK_SMALLER` 
  - `CASE:LIMIT_BID_BIGGER` 
  - `CASE:LIMIT_BID_EQUAL` 
  - `CASE:LIMIT_BID_SMALLER` 
  - `CASE:MARKET_BID_BIGGER` 
  - `CASE:MARKET_BID_EQUAL` 
  - `CASE:MARKET_BID_SMALLER` 
  - `CASE:MARKET_ASK_BIGGER` 
  - `CASE:MARKET_ASK_EQUAL` 
  - `CASE:MARKET_ASK_SMALLER`
- `STREAM:PLACEMENT`: 주문이 오더북에 등록되었을 때 발생하는 stream 입니다.
- `STREAM:REFUND`: 주문에 대한 환불 처리가 필요한 경우 발생하는 stream 입니다.
- `STREAM:ERROR`: 프로세스에 에러가 발생한 경우, 로깅을 위해 발생하는 stream 입니다.
