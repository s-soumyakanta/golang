package main

import (
	"fmt"
	"time"
)

//Structs are custom data structures in go. These are like similar to classes of other language

// Struct Embedding - to do this wee need to pass it to order struct simly
type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nono second precision
	customer
}

// CONSTRUCTOR hack
func newOrder(id string, amount float32, status string) *order {
	//initial setup goes here
	allNewOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &allNewOrder
}

// o is the receiver type. We use o here because it generally used convection as per first letter of struct

// We are modifying the value so need the pointer
func (o *order) changeStatus(status string) {
	//struct automatically doing the derefferncing we won't need to add *
	o.status = status
}

// We are just getting the value so no need for pointer
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// If we do not set any value then it is empty value. It will store 0 and empty string for int and string accordingly & false for bool.
	recentOrder := order{
		id:     "23",
		amount: 45,
		status: "In transit",
	}

	fmt.Println(recentOrder) //{23 45 In transit {0 0 <nil>}}
	recentOrder.createdAt = time.Now()
	fmt.Println(recentOrder) //{23 45 In transit {13954861578011495512 515201 0xe48b20}}

	//Printing any elem
	fmt.Println(recentOrder.status) //In transit

	myOrder := order{
		id:        "9",
		amount:    800,
		status:    "Received",
		createdAt: time.Now(),
	}
	fmt.Println(myOrder) //{9 800 Received {13954861767205346936 1683301 0xf48b20}}

	myOrder.changeStatus("Delivered")
	fmt.Println(myOrder.status) //Delivered

	fmt.Println(myOrder.getAmount()) //800

	brandNewOrder := newOrder("3", 888, "Initiated")
	fmt.Println(brandNewOrder) //&{3 888 Initiated {0 0 <nil>}}

	//If we are using struct for once only - Inline structs

	language := struct {
		name     string
		isTrendy bool
	}{"Golang", true}

	fmt.Println(language) //{Golang true}

	//struct embedding
	// newCustomer := customer{
	// 	name:  "S S",
	// 	phone: "3432432",
	// }
	orderWithCustomer := order{
		id:     "3",
		amount: 34,
		status: "received",
		// customer: newCustomer,
		customer: customer{
			name:  "Soumya",
			phone: "3333",
		},
	}

	// fmt.Println(orderWithCustomer) //{3 34 received {0 0 <nil>} { }} - empty for customer
	fmt.Println(orderWithCustomer) //{3 34 received {0 0 <nil>} {Soumya 3333}}

	orderWithCustomer.customer.name = "S Soumyakanta"
	fmt.Println(orderWithCustomer.customer.name) //S Soumyakanta
}
