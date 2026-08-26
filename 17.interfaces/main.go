package main

import "fmt"


// interfaces

type paymenter interface {
	pay(amount float32)
}


// payment method 
type payment struct {
	Gw paymenter
}

func (p payment) makepayment(amount float32) {

	p.Gw.pay(amount)
}




// razorpay
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("payment done with razorpay :- ", amount)
}


// payfast 
type payfast struct{}

func (p payfast) pay(amount float32) {
	fmt.Println("payment done with payfast :- ", amount)
}


// fake payment gateway


type fakepayment struct{}

func (f fakepayment) pay(amount float32){
	fmt.Println("payment done with fakepaymentgw :- ", amount)
}


func main() {

	f := fakepayment{}
	// r := razorpay{}
	// p := payfast{}


	newpay := payment{Gw: f}

	newpay.makepayment(50)
}
