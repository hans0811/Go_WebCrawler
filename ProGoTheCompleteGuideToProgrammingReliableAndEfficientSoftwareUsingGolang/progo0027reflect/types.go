package main

import "fmt"

func Printfl1n(template string, values ...interface{}){
	fmt.Printf(template + "\n", values...)
}

type Product struct {
	Name, Category string
	Price          float64
}

type Customer struct{
	Name, City string
}