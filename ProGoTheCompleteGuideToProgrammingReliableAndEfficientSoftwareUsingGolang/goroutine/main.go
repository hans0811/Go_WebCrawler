package main

import (
	"fmt"
	"time"
)


func receiveDispatches(channel <-chan DispatchNotification) {
    for details := range channel {
        fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
            "x", details.Product.Name)
    }
    fmt.Println("Channel has been closed")
}

func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList {
        select {
            case channel <- p:
                fmt.Println("Sent product:", p.Name)
            default:
                fmt.Println("Discarding product:", p.Name)
                time.Sleep(time.Second)
        }
    }
    close(channel)
}


func main() {
    // fmt.Println("main function started")
    // CalcStoreTotal(Products)
	// time.Sleep(time.Second * 5)
    // fmt.Println("main function complete")


	//Sending and Receiving an Unknown Number of Values

	dispatchChannel := make(chan DispatchNotification, 100)

	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
    // var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel
	

	//go DispatchOrders(dispatchChannel)

	// for {
	// 	if details, open := <- dispatchChannel; open {
    //         fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
    //             "x", details.Product.Name)
    //     } else {
    //         fmt.Println("Channel has been closed")
    //         break
    //     }
	// }

	// for details := range dispatchChannel {

	// 	fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
	// 	"x", details.Product.Name)
	// }

	// fmt.Println("Channel has been closed")

	go DispatchOrders(chan<- DispatchNotification(dispatchChannel))
    //receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
	productChannel := make(chan *Product, 5)
	go enumerateProducts(productChannel)
    openChannels := 2

	for {
        select {
            case details, ok := <- dispatchChannel:
                if ok {
                    fmt.Println("Dispatch to", details.Customer, ":",
                        details.Quantity, "x", details.Product.Name)
                } else {
                    fmt.Println("Channel has been closed")
                    dispatchChannel = nil
                    openChannels--
                }
			case product, ok := <- productChannel:
                if ok {
                    fmt.Println("Product:", product.Name)
                } else {
                    fmt.Println("Product channel has been closed")
                    productChannel = nil
                    openChannels--
                }
            default:
                if (openChannels == 0) {
                    goto alldone
                }
                fmt.Println("-- No message ready to be received")
                time.Sleep(time.Millisecond * 500)
        }
    }
    alldone: fmt.Println("All values received")


}