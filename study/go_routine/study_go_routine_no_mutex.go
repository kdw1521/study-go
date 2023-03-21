package goroutine

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업완료 - 결과 %d\n", j.index, j.index*j.index)
}

func Job_Runner(wg *sync.WaitGroup) {
	var jobList [10]Job
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
