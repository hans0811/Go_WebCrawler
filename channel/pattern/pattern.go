package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s: message %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleanup done")
				done <- struct{}{}
				return
			}
			i++
		}

	}()

	return c
}

func fanIn(chs ...chan string) chan string {
	c := make(chan string)

	for _, ch := range chs {

		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)

	}

	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)

	go func() {
		for {

			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

func nonBlockingWaiting(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}

}

func main() {
	done := make(chan struct{})
	//m1 := msgGen("service1") // publish message, handle
	//m2 := msgGen("service2")

	//	for {
	//fmt.Println(<-m1)
	//fmt.Println(<-m2)

	//fmt.Println(<-m1)
	//
	//if m, ok := timeoutWait(m1, 2*time.Second); ok {
	//	fmt.Println(m)
	//} else {
	//	fmt.Println("timeout")
	//}

	//if m, ok := nonBlockingWaiting(m2); ok {
	//	fmt.Println(m)
	//} else {
	//	fmt.Println("no message from service2")
	//}
	//	}

	m1 := msgGen("service1", done) // publish message, handle
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWait(m1, time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
		}
	}

	done <- struct{}{}
	<-done

}
