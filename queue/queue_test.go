package queue

import "fmt"

func ExampleQueue_Pope() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	//1
	//2
	//false

}
