package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {

	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 8, 10, 8}
	var grid [4][5]int

	fmt.Println("arr1")
	printArray(&arr1)
	fmt.Println("arr3")
	printArray(&arr3)
	fmt.Println("All")
	fmt.Println(arr1, arr3)
	//printArray(arr2)

	fmt.Println("")
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	maxi := -1
	maxValue := -1
	for i, v := range arr3 {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}

	fmt.Println(maxi, maxValue)
}
