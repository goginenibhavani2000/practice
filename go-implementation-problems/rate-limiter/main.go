package main

/*
You're building a service that calls an external API (say, an LLM provider). The API allows you to make at most 5 requests per second. If you exceed that, they return a 429 error and your account gets throttled.
But your traffic is bursty — sometimes you go 10 seconds with no requests, then suddenly need to make 8 requests at once. You'd like to handle short bursts without hitting the limit, as long as you stay within the rate over time.
Write a RateLimiter type that decides whether each request is allowed.
*/
import (
	"sync"
	"time"
)

type RateLimiter struct {
	ratePerSec int
	burst      int
	tokensLeft float64 // how many tokens left in the bucket
	lastRefill time.Time
	mu         sync.Mutex
}

func NewRateLimiter(ratePerSec int, burst int) *RateLimiter {
	return &RateLimiter{ratePerSec: ratePerSec, burst: burst, tokensLeft: float64(burst), lastRefill: time.Now()}
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now()
	// how much time has passed since last refill
	elapsed := now.Sub(r.lastRefill)
	//  how many tokens can be added to the bucket
	tokensadded := elapsed.Seconds() * float64(r.ratePerSec)
	// if token that can be added(ex-9.0) with token left(ex-2.0) is mmore than burst(ex-10.0), limit tokens to burst(10.0)
	if tokensadded+r.tokensLeft > float64(r.burst) {
		r.tokensLeft = float64(r.burst)
	} else {
		r.tokensLeft += tokensadded
	}
	// update last refill to now
	r.lastRefill = now
	// if for each allow call atleast 1 token is left, decrease tokenleft and return that the request is allowed
	if r.tokensLeft >= 1 {
		r.tokensLeft -= 1
		return true
	}

	return false
}
