package main

import "fmt"
func CalcStoreTotal(data ProductData) {
    var storeTotal float64
    for category, group := range data {
        storeTotal += group.TotalPrice(category)
    }
    fmt.Println("Total:", ToCurrency(storeTotal))
}
func (group ProductGroup) TotalPrice(category string, ) (total float64) {
    for _, p := range group {
		fmt.Println(category, "product:", p.Name)
        total += p.Price
    }
    fmt.Println(category, "subtotal:", ToCurrency(total))
    return
}