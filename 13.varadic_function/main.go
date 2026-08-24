package main

import "fmt"

func sum(a ...int) int {
	total := 0

	for _, nums := range a {
		total += nums
	}

	return total
}

func main() {

	result := sum(2, 3, 4, 5, 6, 7, 8, 9, 0, 12, 12, 13, 14, 45, 56, 67)

	nums := []int{3,4,5,6}


	fmt.Println(result)

	resul := sum(nums...)
	




	fmt.Println(resul)

}
