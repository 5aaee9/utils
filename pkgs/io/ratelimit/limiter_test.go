package ratelimit

import (
	"testing"
	"time"
)

func TestLimiterTryNonBlock(t *testing.T) {
	limit := 1024
	limiter := NewLimiter(uint64(limit))
	start := time.Now()

	for i := 0; i < limit; i++ {
		limiter.Wait(1)
	}

	if time.Now().Sub(start) >= time.Second {
		t.Error("The limiter blocked when it shouldn't have")
	}
}


func TestLimiterTryWithBlock(t *testing.T) {
	limit := 1024
	limiter := NewLimiter(uint64(limit))
	start := time.Now()

	for i := 0; i < limit; i++ {
		limiter.Wait(1)
	}

	limiter.Wait(1)

	if time.Now().Sub(start) <= time.Second {
		t.Error("The limiter didn't block when it should have")
	}
}