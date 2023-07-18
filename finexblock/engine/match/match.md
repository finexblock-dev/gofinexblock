# Match

## Description

`STREAM:MATCH`의 Consumer.

매칭 스트림을 전달받고, 매칭에 따른 체결 이벤트를 생성 및 전달한다.

```go
package _

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/engine/match"
	"github.com/redis/go-redis/v9"
)

func main() {

	var cluster *redis.ClusterClient
	
	engine := match.New(cluster)
	
	go engine.Consume()
	go engine.Claim()
	
	// ...
}
```