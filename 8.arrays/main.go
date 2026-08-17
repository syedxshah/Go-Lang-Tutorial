package main

import "fmt"

func main() {


// simple arrays 
// arrays are fixed size data structure

	var a [5]int
	var b [90]int
	fmt.Println(len(a))
	fmt.Println(len(b))

	fmt.Println("elements in the b arrays ")
	b[0] = 1
	fmt.Println(b)


	 var c  [3]interface{}
	c[0] = nil
	c[1] = 12
	c[2] = "hello"
	fmt.Println(c) 


	// 2d arrays
	d := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(d)

}
