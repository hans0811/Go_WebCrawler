package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product{
	return &Product{name, category, price}
}

func (p *Product) Price(taxRate float64) float64{
	return p.price + (p.price * taxRate)
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


// Understanding When Promotion Cannot Be Performed

type SpecialDeal struct{
	Name string
	*Product
	price float64
}

func NewSpecialDeal(name string, p *Product, discount float64) *SpecialDeal{
	return &SpecialDeal{name, p, p.price-discount}
}

func (deal *SpecialDeal) GetDetails() (string, float64, float64){
	return deal.Name, deal.price, deal.Price(0)
}

// Create new Method rewrite Price
func (deal *SpecialDeal) Price(taxRate float64) float64{
	return deal.price
}

// Understanding Promotion Ambiguity



// func flatten(root *TreeNode)  {
    
//     //node := &root;

// 	for root != nil{

// 		left := root.Left
// 		right := root.Right

// 		if( left != nil){
// 			temp := left;
// 			for temp.Right != nil{
// 				temp = temp.Right
// 			}
// 			temp.Right = right
// 			root.Right = left
// 			root.Left = nil
// 		}
//         //fmt.Println(node.Val)
// 		root = root.Right
// 	}
// }

func flatten(root *TreeNode)  {
	t := root

	for root != nil{

		left := root.Left
		right := root.Right

		if( left != nil){
			temp := left;
			for temp.Right != nil{
				temp = temp.Right
			}
			temp.Right = right
			//node.Right = left
			//node.Left = nil
		}
		fmt.Println(t)
		//node = node.Right
	}
}

func isAnagram(s string, t string) bool {
    
	dict := [26]int{}

	for i := 0; i < len(s); i++{
		idx := s[i] - 'a'
		dict[idx]++
	}


	for i := 0; i < len(t); i++{
		idx := t[i] - 'a'
		dict[idx]--
	}

	for i := 0; i < len(dict); i++{

		if dict[i] != 0 {
			return false
		}
	}

	return true;
    
}

//Creating a Chain of Nested Types

//Using Multiple Nested Types in the Same Struct
type Crew struct{
	Captain, FirstOfficer string
}

type RentalBoat struct{
	*Boat
	IncludeCrew bool
	*Crew
}

func NewRentalBoat( name string, price float64, capacity int, 
	motorized, crewed bool, 
	captain, firstOfficer string ) *RentalBoat {
	return &RentalBoat{
		NewBoat(name, price, capacity, motorized), 
		crewed,
		&Crew{captain, firstOfficer}}
}



func main() {

	
	// b := TreeNode{2, nil, nil}
	// c := TreeNode{5, nil, nil}

	// a := TreeNode{1, &b, &c}

	// flatten(&a)

	// fmt.Println(a.val)
	//fmt.Println("anv")

	//isAnagram("a","ab")


	//kayak := store.NewProduct("Kayak", "Watersports", 275)
	// lifejacket := &store.Product{ Name: "Lifejacket", Category:  "Watersports"}

	// for _, p := range []*store.Product { kayak, lifejacket}{
	// 	fmt.Println("Name:", p.Name, 
	// 				"Category:", p.Category,
	// 				"Price:", p.Price(0.2))	
	// }


	// boats := []*Boat {
    //     NewBoat("Kayak", 275, 1, false),
    //     NewBoat("Canoe", 400, 3, false),
    //     NewBoat("Tender", 650.25, 2, true),
    // }
    // for _, b := range boats {
    //     fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name, "Price:", b.Price(0.2) )
    // }

	// boat := NewBoat("Kayak1", 275, 1, false)

	// boat.Name = "Green Kayak"

	// Error
	// boat1 := Boat { Name: "Kayak", Category: "Watersports",
    // Capacity: 1, Motorized: false }

	boat2 := Boat { Product: &Product{ Name: "Kayak",
    Category: "Watersports"}, Capacity: 1, Motorized: false }

	boat2.Name = "Green Kayak2"

	// Creating a Chain of Nested Types
	// rentals := []*RentalBoat {
    //     NewRentalBoat("Rubber Ring", 10, 1, false, false),
    //     NewRentalBoat("Yacht", 50000, 5, true, true),
    //     NewRentalBoat("Super Yacht", 100000, 15, true, true),
    // }
    // for _, r := range rentals {
    //     fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2))
    // }

	// rentals := []*RentalBoat {
    //     NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
    //     NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
    //     NewRentalBoat("Super Yacht", 100000, 15, true, true,"Dora", "Charlie"),
    // }

    // for _, r := range rentals {
    //     fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2),
    //         "Captain:", r.Captain)
    // }
	// product := NewProduct("Kayak", "Watersports", 279)

	// deal := NewSpecialDeal("Weekend Special", product, 50)
    // Name, price, Price := deal.GetDetails()
    // fmt.Println("Name:", Name)
    // fmt.Println("Price field:", price)
    // fmt.Println("Price method:", Price)


	//Promotion Ambiguity
	// kayak4 := NewProduct("Kayak", "Watersports", 279)

    // type OfferBundle struct {
    //     *SpecialDeal
    //     *Product
    // }
    // bundle := OfferBundle {
    //     NewSpecialDeal("Weekend Special", kayak4, 50),
    //     NewProduct("Lifrejacket", "Watersports", 48.95),
    // }
    // fmt.Println("Price:", bundle.Price(0))
	
	
	//Using Composition to Implement Interfaces

	// products := map[string] store.ItemForSale{
	// "Kayak" : store.NewBoat("Kayak", 279, 1, false),
	// "Ball" : store.NewProduct("Soccer Ball", "Soccer", 19.50),
	// }

	// for key, p := range products{
	// 	//fmt.Println("Key:", key, "Price:", p.Price(0.2))

	// 	switch item := p.(type) {
	// 	case *store.Product:
	// 		fmt.Println("Name:", item.Name, 
	// 					"Category:", item.Category,
	// 					"Price:", item.Price(0.2))

	// 	case *store.Boat:
	// 		fmt.Println("Name:", item.Name, 
	// 					"Category:", item.Category,
	// 					"Price:", item.Price(0.2))

	// 	default:
	// 		fmt.Println("Key:", key, "Price:", p.Price(0.2))
	// 	}
	// }

	// for key, p := range products{

	// 	switch item := p.(type) {
	// 	case store.Describalbe:
	// 		fmt.Println("Name:", item.GetName(), 
	// 					"Category:", item.GetCategory(),
	// 					"Price:", item.Price(0.2))


	// 	default:
	// 		fmt.Println("Key:", key, "Price:", p.Price(0.2))
	// 	}
	// }


	first := 100
	second := first
	first++

	fmt.Println(first, second)

	a := 100
	b := &a // var b *int = &a
	a++
	*b++
	fmt.Println(a, *b)

	var myPointer *int
	myPointer = b
	*myPointer++
	fmt.Println(a, *b) //103 103

	// Pointer need to point value

	//Pointing at Pointers
	x := 100
	y := &x
	z := &y
	fmt.Println(x, *y, **z)

	// Understanding Why Pointers Are Useful
	names := [3]string {"Alice", "Charlie", "Bob"}
	secondName := &names[1]
	fmt.Println(*secondName)
	sort.Strings(names[:])
	fmt.Println(*secondName)

}