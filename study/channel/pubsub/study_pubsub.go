package pubsub

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PubSubRunner() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	publisher := NewPublisher(ctx)
	subscriber1 := NewSubscriber("Wando", ctx)
	subscriber2 := NewSubscriber("Minju", ctx)

	go publisher.Update()

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)

	go subscriber1.Update()
	go subscriber2.Update()

	go func() {
		tick := time.Tick(time.Second * 2)
		for {
			select {
			case <-tick:
				publisher.Publish("Hello MSG")
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()

	fmt.Scanln() // 입력이 들어오면 종료
	cancel()

	wg.Wait()
}
