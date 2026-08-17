package main

import (
	"fmt"
	"time"
)

func main() {

	i := 9
	// simple switch expression 

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	// multi switch

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")
	default:
		fmt.Println("its work day")
	}

	// type switch 

	firstfunc := func(i interface{}){
		switch i.(type){
		case int:
			fmt.Println("its integer")
		case bool:
			fmt.Println("its boolean")
		case float64:
			fmt.Println("its float 64")
		case string:
			fmt.Println("its string")
		case nil:
			fmt.Println("its null")
		default:
			fmt.Println("other", i)
		}
	}


	firstfunc("hello")
	firstfunc(12)
	firstfunc(12.5)
	firstfunc(true)
	firstfunc(nil)
	

}
