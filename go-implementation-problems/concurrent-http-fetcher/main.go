package main

import (
	"net/http"
	"sync"
)

type Result struct {
	URL        string
	StatusCode int
	Err        error
}

func FetchAll(urls []string, maxConcurrent int) []Result {
	results := make(chan Result, len(urls))
	sem := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			sem <- struct{}{}        // ACQUIRE: send into channel (blocks if buffer full)
			defer func() { <-sem }() // RELEASE: receive from channel (frees a slot)
			resp, err := http.Get(url)
			if err != nil {
				results <- Result{URL: url, Err: err}
				return
			}
			defer resp.Body.Close() // safe — only runs if err was nil
			results <- Result{URL: url, StatusCode: resp.StatusCode}
		}(url)
	}

	wg.Wait()
	res := make([]Result, len(urls))
	for i := 0; i < len(urls); i++ {
		res[i] = <-results
	}
	close(results)

	return res
}
