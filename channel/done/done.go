package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		//n := <-c
		//fmt.Println(n)
		fmt.Printf("Worker %d received %c\n",
			id, <-c)
	}
}

// channel can be return
func createWorkerReturnChan(id int) chan<- int {
	c := make(chan int)

	// need use goroutine or it can be deadlock

	go func() {
		for {
			fmt.Printf("Worker %d received % c\n",
				id, <-c)
		}
	}()
	return c
}

// buffered
func workerBuffer(id int, c chan int) {
	for n := range c {
		// check channel is close
		//n, ok := <-c
		//if !ok {
		//	break
		//}

		fmt.Printf("Worker %d received % c\n",
			id, n)
	}
}

func createBufferedWorkerReturnChan(id int) chan<- int {
	c := make(chan int)
	go workerBuffer(id, c)
	return c
}

func chanDemo() {
	// var c chan int // c == nil

	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorkerReturnChan(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		//n := <-channels[i] // cannot br received
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	//var channels [10]chan int
	//for i := 0; i < 10; i++ {
	//	channels[i] = make(chan int)
	//	go worker(i, channels[i])
	//}
	//
	//for i := 0; i < 10; i++ {
	//	channels[i] <- 'a' + i
	//}
	//
	//for i := 0; i < 10; i++ {
	//	channels[i] <- 'A' + i
	//}

	//c := make(chan int)

	// need a goroutine to listen
	//go worker(0, c)

	//go func() {
	//	for {
	//		n := <-c
	//		fmt.Println(n)
	//	}
	//}()

	//c <- 1
	//c <- 2
	//n := <-c
	//fmt.Println(n)
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go workerBuffer(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go workerBuffer(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
