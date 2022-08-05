package store

import "fmt"

func ab() {
	fmt.Sscanln("hello")
}

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product{
	return &Product{name, category, price}
}

func (p *Product) Price(taxRate float64) float64{
	fmt.Print("Product-Price!!!!!!!!!!!!")
	return p.price + (p.price * taxRate)
}

type ItemForSale interface{
	Price(taxRate float64) float64
}

// Composing Types
type Boat struct{
	*Product
	Capacity int
	Motorized bool
}

func NewBoat(name string, price float64, capacity int, motorized bool) *Boat{
	return &Boat {
        NewProduct(name, "Watersports", price), capacity, motorized,
    }
}

//implemented for the *Product type
type Describalbe interface{
	GetName() string
	GetCategory() string
	ItemForSale
}

func (p *Product) GetName() string{
	return p.Name
}

func (p *Product) GetCategory() string{
	return p.Category
}