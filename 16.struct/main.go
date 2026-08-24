package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time // nano second precision
	customer  customer
}

func (o *order) changeId(id string) {
	o.id = id
}

func main() {

	lang := struct {
		name   string
		isgood bool
	}{
		"golang", true}

	fmt.Println(lang)
	my := customer{
		name:  "Ansar",
		phone: "03240745824",
	}
	order := order{
		id:       "a1",
		amount:   50.00,
		status:   "Available",
		customer: my,
	}

	order.createdAt = time.Now()
	order.changeId("b2")
	fmt.Println(order)
}
