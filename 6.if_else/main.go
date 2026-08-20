package main

import "fmt"

func main() {

	age := 12

	if age > 11 {
		fmt.Println("Age Is Greater Than 11")
	}

	check := true 
	
	if check == true {
		fmt.Println(true)
	} else{
		fmt.Println(false)
	}


	name := "ansar"

	if name == "syed" {
		fmt.Println("Syed")
	} else if name == "ansar" {
		fmt.Println("Ansar")
	} else {
		fmt.Println("Ali")

	}

}