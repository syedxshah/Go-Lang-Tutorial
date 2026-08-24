package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multi() (string, string, string) {
	return "syed", "ansar", "ali"
}

// function pass into function

func processit(fn func(a int) int) int {
	return fn(1)
}

// function return into function
func test() func(a int) int {
	return func(a int) int { return a }

}

func main() {
	z := add(3, 5)

	fmt.Println(z)
	v1, v2, v3 := multi()

	fmt.Println(v1, v2, v3)

	fmt.Println(multi())
	f := func(a int) int {
		return a
	}
	fmt.Println(processit(f))

	fi := test()

	fmt.Println(fi(5))

}
