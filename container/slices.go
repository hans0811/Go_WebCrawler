package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100

}

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s := arr[2:6]
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] = ", s1)
	s2 := arr[:]
	fmt.Println("arr[:] = ", s2)

	fmt.Println("After updateDSlice S1")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateDSlice S2")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	//fmt.Println(s1[4]) // cant get
	//fmt.Println("s1=", s1)
	//fmt.Println("s2=", s2)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d",
		s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("")
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	// s4 and s5 no longer view arr
	fmt.Println(arr)

}
