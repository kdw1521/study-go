package main

import (
	// "fmt"
	// "math/rand"
	"study-go/study/go_routine"
	"sync"
	// "time"
	// "time"
)

var wg sync.WaitGroup

func main() {
	// go goroutine.PrintHangul()
	// go goroutine.PringNumber()

	// time.Sleep(3 * time.Second)

	// goroutine.WG.Add(10) // 10개의 작업을 만들어주고
	// for i := 0; i < 10; i++ {
	// 	go goroutine.SumAtoB(1, 1000000000)
	// }
	// goroutine.WG.Wait()

	// var wg sync.WaitGroup

	// account := &goroutine.Account{Balance: 10}
	// wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		for {
	// 			goroutine.DepositAndWithdraw(account)
	// 		}
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// goroutine.Runner(&wg)

	goroutine.Job_Runner(&wg)
}
