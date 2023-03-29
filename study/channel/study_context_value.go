package channel

import (
	"context"
	"fmt"
	"sync"
)

var wg3 sync.WaitGroup

func ContextValueRunner() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9) // number라는 키에 9라는 value를 지정한다.
	go square5(ctx)

	wg.Wait()
}

func square5(ctx context.Context) {
	if v := ctx.Value("number"); v != nil { // 컨텍스트의 값을 가져온다.
		n := v.(int)
		fmt.Printf("Square5:%d", n*n)
	}
	wg.Done()
}
