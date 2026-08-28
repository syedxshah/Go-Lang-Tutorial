package main

import (
	"fmt"
	"sync"
)

func task(id int) {
	fmt.Println("task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		// for j := 0; j < 10; j++ {
		// 	go task(i)
		// }
		// fmt.Println("2nd", i)
		// go func() {
		// 	fmt.Println(i)
		// }()
		// fmt.Println("3rd", i)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(i)
		}(i, &wg)
	}
	wg.Wait()
}
