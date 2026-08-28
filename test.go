package main

import (
	"fmt"
)

type null struct{}

func describe[T bool | int | string | float64 | float32 | null](t T) {

	switch v := any(t).(type) {

	case int:
		fmt.Println("int", test(v))
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	case float64:
		fmt.Println("float64", v)
	case float32:
		fmt.Println("float32", v)
	case null:
		fmt.Println("null")
	default:
		fmt.Println("unknown type")
	}

}

func test(a int) string {

	if a%2 == 0 {
		return "even"
	}
	return "odd"

}

func main() {

	describe(null{})

	describe(10)

	describe("hello")

	describe(true)

	describe(3.14)

	describe(783)

}
