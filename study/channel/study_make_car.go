package channel

import (
	"fmt"
	"sync"
	"time"
)

type Car struct { // 차 구조체 생성
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func MakeCar() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go MakeTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

func MakeBody(tireCh chan *Car) { // 차체 생성
	tick := time.Tick(time.Second)        // 1초마다 생성
	after := time.After(10 * time.Second) // 10초가 지나면 종료

	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Wando Car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func MakeTire(tireCh, paintCh chan *Car) { // 바퀴 생성
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Wando tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) { // 페인트 칠
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime) // 경과 시간 출력
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
