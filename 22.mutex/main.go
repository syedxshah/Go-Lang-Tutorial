package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()

		wg.Done()
	}()

	p.mu.Lock()
	p.views++

}

func (p *post) print() {

	fmt.Println("Total Views Is :- ", p.views)
}
func main() {

	p := post{views: 0}
	var wg = sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go p.inc(&wg)
	}

	wg.Wait()
	p.print()

}
