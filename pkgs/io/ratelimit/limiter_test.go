package ratelimit

import (
	"sync"
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

func TestLimiterTryNonBlockOnZeroValue(t *testing.T) {
	limit := 1024
	limiter := NewLimiter(0)
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

func TestLimiterMutliThread(t *testing.T) {
	limit := 1024
	limiter := NewLimiter(uint64(limit))
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(limit*5 + 1)

	for i := 0; i < (limit*5 + 1); i++ {
		go func() {
			limiter.Wait(1)
			wg.Done()
		}()
	}

	wg.Wait()

	if time.Now().Sub(start) <= time.Second*5 {
		t.Error("The limiter didn't block when it should have")
	}
}
