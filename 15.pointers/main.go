package main

import "fmt"

// by value
func chnageNUM (num  int ){
	num = 5
	fmt.Println("In chnageNUM",num) 
}


// by reference
func chnageNUm (num  *int ){
	*num = 5
	fmt.Println("In chnageNUM",*num) 
}



func main()  {

	num:= 1 

	// chnageNUM(num)

	chnageNUm(&num)

	fmt.Println("after Changenum value is  :",num)
	
}
