package main

import (
	"001_go_env/queue"
	"fmt"
)

func main() {

	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Print(q.Pop())

}
