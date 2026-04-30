package main

import (
	"fmt"
	"sync"
)

/*
Fan-out: square 100 numbers using 5 workers

Input: 1..100, output: 1, 4, 9, ..., 10000
5 worker goroutines reading from a jobs channel, writing to a results channel
Use sync.WaitGroup to know when workers are done
Refreshes: channel direction (<-chan, chan<-), the worker pattern, close(channel).
*/

/*
You used the canonical pattern. The wg.Wait() → close(results) goroutine is the textbook way to handle "multiple writers, one reader." This is the part you got wrong on the previous attempt and absorbed in one round.
3. main reads the results directly. You removed the unnecessary reader goroutine. That's the mature choice — it eliminates the race and removes a goroutine you didn't need.
4. You used defer close(jobs) on the producer. That's the right pattern: the moment the producer's loop exits, the channel closes. Even if the loop panics mid-way, defer still fires.
5. You're now thinking about channel ownership. Workers don't close results (they're not the sole writers). The producer closes jobs (it's the sole writer). The closer goroutine closes results after all writers are done.
*/

func main() {
	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range jobs {
				results <- num * num
			}
		}()
	}

	go func() {
		defer close(jobs)
		for i := 1; i <= 100; i++ {
			jobs <- i
		}
	}()
	go func() {
		wg.Wait()      // wait for all 5 workers to finish
		close(results) // safe to close — no more writers
	}()
	for ans := range results {
		fmt.Println(ans)
	}
}
