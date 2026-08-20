package main

import "fmt"

var hight float64 = 5.9

func main() {

	const name = "John"

	const age = 30
	// age = 15

	const isStudent = true

	const (
		port    = 8080
		host    = "localhost"
		country = "Pakistan"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isStudent)
	fmt.Println(port)
	fmt.Println(host)
	fmt.Println(country)
	fmt.Println(hight)
}
