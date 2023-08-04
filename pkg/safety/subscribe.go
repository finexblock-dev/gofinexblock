package safety

import "log"

type Subscriber interface {
	Subscribe()
}

func InfinitySubscribe(subscriber Subscriber) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("recovered from panic : %v", r)
				}
			}()
			subscriber.Subscribe()
		}()
	}
}