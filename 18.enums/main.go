	package main

	import "fmt"

	type orderstatus int

	const (
		Recieved orderstatus = iota
		Conform
		Prepared
		Deliver
	)

	func check(o orderstatus) string {
		switch o {
		case 0:
			return "recived"
		case 1:
			return "conform"
		case 2:
			return "prepared"
		case 3:
			return "deliver"
		}
		return "Error"
	}

	func checkorderstatus(o orderstatus) {

		status := check(o)
		fmt.Println(" the order status is ",status," and the status code is ",o)
	}

	func main() {

		checkorderstatus(Prepared)

	}
