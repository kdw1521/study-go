package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func ContextRunner() {
	wg2.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) // 기본 컨텍스트 하나 반환
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	wg2.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg2.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
