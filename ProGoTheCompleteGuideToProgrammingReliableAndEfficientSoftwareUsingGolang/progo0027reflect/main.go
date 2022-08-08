package main

import (
	"fmt"
	"reflect"
	"strings"
)

func ppp(){
	fmt.Print()
}

func printDetailsRaw(values ...Product) {
	for _, elem := range values {

		Printfln("Product: Name : %v, Category: %v, Price: %v", elem.Name, elem.Category, elem.Price)
	}
}

func printDetails(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Struct{
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails, fmt.Sprintf("%v: %v", fieldName, fieldVal))
			}
			Printfln(" %v:%v", elemType.Name(), strings.Join(fieldDetails, ", "))
		}else{
			Printfl1n("%v: %v", elemType.Elem(), elemValue)
		}
	}
}

type Payment struct{
	Currency string
	Amount float64
}

func main(){
	product := Product {
		Name: "Kayak", Category: "Watersports", Price: 279,
	}
	customer := Customer { Name: "Alice", City: "New York" }
	payment := Payment { Currency: "USD", Amount: 100.50 }

	printDetails(product, customer, payment)
	//printDetailsByTypes(product, customer)
}