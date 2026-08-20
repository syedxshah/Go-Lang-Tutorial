package main

import "fmt"

func main() {
	var name = "John"

	var age = 30

	var isStudent = true

	var (
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

	fmt.Println(isStudent)

	var hight float64

	hight = 5.9

	fmt.Println(hight)

	// shorthand variable declaration
	porti := 8080
	hosti := "localhost"
	countryi := "Pakistan"

	fmt.Println(porti)
	fmt.Println(hosti)
	fmt.Println(countryi)
}
