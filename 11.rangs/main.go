package main

import (
	"fmt"
	
)

func main () {

	s := []int{6,7,8}

	for i , val :=range s {

		fmt.Println(i ,  val)
	}

	m := map[string]interface{}{"fname" : "ansar" , "lnmae" : "ali" } 

	for k , val := range m {
		fmt.Println(k,val)
	}

	for k  := range m {
		fmt.Println(k)
	}


	// its uni code not hex. code -> if we use the range on the string it give decimal values
	// we use use the string(c) to convert it .....

	for i , c := range "syed ansali ali aftab " {
		fmt.Println(i, c)

		fmt.Println(i, string(c))
	}



}
