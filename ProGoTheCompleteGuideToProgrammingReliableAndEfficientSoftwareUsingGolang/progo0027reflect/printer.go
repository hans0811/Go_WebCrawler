package main

import "fmt"

func Printfln(template string, values ...interface{}){
	fmt.Printf(template + "\n", values...)
}

// use empty interface and switch
func printDetailsByTypes(values ...interface{}) {
	for _, elem := range values {
	switch val := elem.(type) {
	case Product:
	Printfln("Product: Name: %v, Category: %v,Price: %v",
	val.Name, val.Category, val.Price)
	case Customer:
	Printfln("Customer: Name: %v, City: %v",
	val.Name, val.City)
	}
	}
}