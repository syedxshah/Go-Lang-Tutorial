package main

import (
	"fmt"
	"maps"
)



func main() {
	m := make(map[string]interface{})

	 test :=  make(map[string]map[string]interface{})

	m["name"] = "ansar"
	m["age"] = 20

	v , ok := m["name"]

	if ok{
		fmt.Println(v)
	} else{
		fmt.Println("not all")
	}

	test["user1"] = make(map[string]interface{}) 

	test["user1"]["name"] = "ansar"

	fmt.Println(m["name"] , m["age"])

	fmt.Println(test)

	var masp  map[string]interface{}

	fmt.Println(len(masp))

	fmt.Println(test)

	clear(test)

	clear(m)


	fmt.Println(test)

	fmt.Println(maps.Equal(m ,masp))



	

}

