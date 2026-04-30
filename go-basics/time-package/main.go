package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Retry a function with exponential backoff

A function that returns an error 60% of the time (simulate with rand)
Retry up to 3 times, doubling the wait between attempts (100ms, 200ms, 400ms)
Library: time, math/rand, fmt.Errorf for wrapping
Refreshes: time.Sleep, error wrapping with %w, custom retry logic.
*/
func generror(retry int) error {
	if rand.Float64() < 0.6 {
		return fmt.Errorf("retry: %d ,error occured", retry)
	}

	return nil
}
func main() {
	retry := 3
	sleep := 100 * time.Millisecond
	var err error
	for i := 0; i < retry; i++ {
		err = generror(i)
		if err == nil {
			fmt.Printf("Succeded in attempt: %d", i)
			break
		}
		if i != retry-1 {
			time.Sleep(sleep)
			sleep = 2 * sleep
		}
	}
	if err != nil {
		fmt.Print("All attampts failed")
	}
}
