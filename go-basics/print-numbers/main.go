package main

import (
	"fmt"
	"sync"
)

/*
Goroutine A prints odd, goroutine B prints even
Use two channels to coordinate
Refreshes: go keyword, unbuffered channels for synchronization, sync.WaitGroup.
*/

func main() {
	odd := make(chan bool)
	even := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 1; i < 10; i = i + 2 {
			<-odd
			fmt.Printf("goroutineName: A , num: %d\n", i)
			if i < 10 {
				even <- true
			}

		}
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 2; i <= 10; i = i + 2 {
			<-even
			fmt.Printf("goroutineName: B , num: %d\n", i)
			if i < 10 {
				odd <- true
			}
		}

	}(&wg)

	odd <- true
	wg.Wait()
}
