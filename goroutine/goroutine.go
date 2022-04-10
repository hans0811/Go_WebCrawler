package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	var a [10]int
	for i := 0; i < 10; i++ {
		go func(ii int) {

			for {
				a[ii]++
				runtime.Gosched()
				//fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}

	// let main() not execute fast
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
