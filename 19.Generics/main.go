package main

import (
	"fmt"
	
)

func printslice[T any](items []T)  {

	for _, i := range items{

		fmt.Println(i)
	} 
	
}


type stack[t any] struct{
	element []t 
}



func main (){

	s := stack[string]{
		element: []string{"hello","Word","form","our","side"},
	}
	fmt.Println(s)


	slic :=  []string{"ansar","ali","shah"}
	printslice(slic)

}


/// javascript , html ,  css , java , cpp , c , java , golang , dart , python  
 