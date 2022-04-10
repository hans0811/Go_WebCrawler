package main

import (
	"fmt"
	"sync"
)

// buffered
func doWorkerBuffer(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received % c\n",
			id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

// channel can be return
func createWorkerReturnChan(id int, wg *sync.WaitGroup) worker {

	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWorkerBuffer(id, w)
	return w
}

func chanDemo() {
	// var c chan int // c == nil

	var workers [10]worker

	// wait group
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		workers[i] = createWorkerReturnChan(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()

	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}

}

func main() {
	chanDemo()
}
