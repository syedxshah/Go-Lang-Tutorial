package main

import "fmt"

func main() {
	// go only have for loop, while and do while loop is not available in go
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("infinite loop")
		break
	}

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		println(i)
	}

}
