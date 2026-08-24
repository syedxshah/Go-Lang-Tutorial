package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Println("Slice Topic")

	// slice is the empty array/list // dynamic array
	var nums []int

	fmt.Println(nums == nil)

	fmt.Println("length of nums slice", len(nums))

	fmt.Println(cap(nums))

	fmt.Println("By Using Append Function")

	nums = append(nums, 1)

	fmt.Println("length of nums slice", len(nums))
	fmt.Println(cap(nums))
	fmt.Println(nums)

	fmt.Println("By Using For Loop")
	 nums1:=  []int{} 

	for i:= 0 ; i <10 ; i++{
		nums1 = append(nums1, i)
	}
	fmt.Println("length of nums1 slice", len(nums1))
	fmt.Println("capacity of nums1 slice", cap(nums1))
	fmt.Println(nums1)

	fmt.Println(nums1[0:9])

	fmt.Println(slices.Equal(nums , nums1))

}
