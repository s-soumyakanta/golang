package main

import "fmt"

//enumerated types

//To avoid typos
//In inbuild enums in go, we make it using constant and custom types together

//custom types
type myType string
type OrderStatus int

//enums
//iota for intgers
//Received  OrderStatus = "received"
//iota for intgers or only string if we want to store string

const (
	Received  OrderStatus = iota //0
	Confirmed                    //1
	Prepared                     //2
	Deliverd                     //3
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	// changeOrderStatus(Received)
	//Changing order status to Received
	changeOrderStatus(Received)
}
