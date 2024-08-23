package main

import "fmt"

//Interfaces
//naming convection to use er version in last
type paymenter interface {
	pay(amaount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {

	// razorpayPaymentGateWay := razorpay{}
	// razorpayPaymentGateWay.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}

//After Interface we can use any payment gateway service
//Here is one for test

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("Making  payment using fake gateway for testing purpose")
}

// Now we add any new payment gateway without modifing our code
//Let's add paypal

type paypal struct {
}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {

}

//Clean Architecture
//dependency inversion
func main() {
	// razorpayPaymentGw := razorpay{}
	// fakeGw := fakepayment{}
	paypalGw := paypal{}
	newPayment := payment{
		// gateway: razorpayPaymentGw,
		// gateway: fakeGw,
		gateway: paypalGw,
	}
	// newPayment.makePayment(200)
	//Making payment using razorpay 200

	// newPayment.makePayment(200)
	//Making  payment using fake gateway for testing purpose

	newPayment.makePayment(200)
	//Making payment using paypal 200

}
