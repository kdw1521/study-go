package channel

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)
	fmt.Println("Square:", n*n)
	wg.Done()
}

func Runner() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

// chan 크기 관련
func square2() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep!")
	}
}
func Runner2() {
	ch := make(chan int, 1)

	go square2()
	ch <- 9
	fmt.Println("Never print")
}

// 채널에서 대기
func square3(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println("Square3:", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func Runner3() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square3(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
}

// 일정 간격으로 실행 + select
func square4(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("Square:", n*n)
			time.Sleep(time.Second)
		}
	}
}

func Runner4() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square4(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	wg.Wait()
}
