package ratelimit

import (
	"sync"
	"time"
)

// Limiter is counter to limit data
type Limiter struct {
	lock     sync.Mutex
	size     uint64
	now      uint64
	duration time.Duration
	last     time.Time
}

// NewLimiter returns new limiter
func NewLimiter(size uint64) *Limiter {
	return &Limiter{
		size:     size,
		now:      0,
		duration: time.Second,
		last:     time.Now(),
	}
}

// Wait allow you take data
func (l *Limiter) Wait(size uint64) {
	var s uint64
	s = size
	for {
		var ok bool
		var t time.Duration
		ok, s, t = l.Try(s)
		if ok {
			break
		}

		time.Sleep(t)
	}
}

func (l *Limiter) Try(size uint64) (ok bool, remainingCount uint64, waitTime time.Duration) {
	l.lock.Lock()
	defer l.lock.Unlock()

	now := time.Now()
	if now.Sub(l.last) > l.duration {
		l.now = 0
	}

	reamining := l.size - l.now
	if reamining >= size {
		l.now += size
		return true, 0, 0
	}

	l.now += reamining
	return false, size - reamining, l.duration - now.Sub(l.last)
}

// reset internal for test only
func (l *Limiter) reset() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.now = 0
}
