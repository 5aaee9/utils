package ratelimit

import (
	"io"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type dummyReadWriter struct {
	closed atomic.Bool
}

func (r *dummyReadWriter) Close() error {
	r.closed.Store(true)
	return nil
}

func (r *dummyReadWriter) Read(p []byte) (int, error) {
	if r.closed.Load() {
		return 0, io.ErrUnexpectedEOF
	}

	return 1, nil
}

func (r *dummyReadWriter) Write(p []byte) (int, error) {
	if r.closed.Load() {
		return 0, io.ErrUnexpectedEOF
	}

	return 1, nil
}

func TestRateLimitReadWriteCloserRead(t *testing.T) {
	data := make([]byte, 1024)
	limit := 1024
	limited := NewRateLimitReaderWriterCloser(&dummyReadWriter{}, uint64(limit))

	start := time.Now()
	for i := 0; i < limit; i++ {
		_, err := limited.Read(data)
		assert.NoError(t, err)
	}

	if time.Now().Sub(start) >= time.Second {
		t.Error("The limiter blocked when it shouldn't have")
	}

	_, _ = limited.Read(data)

	if time.Now().Sub(start) <= time.Second {
		t.Error("The limiter didn't block when it should have")
	}
}

func TestRateLimitReadWriteCloserWrite(t *testing.T) {
	data := make([]byte, 1024)
	limit := 1024
	limited := NewRateLimitReaderWriterCloser(&dummyReadWriter{}, uint64(limit))

	start := time.Now()

	limited.(*RateLimitReadWriteCloser).limiter.reset()

	for i := 0; i < limit; i++ {
		_, err := limited.Write(data)
		assert.NoError(t, err)
	}

	_, _ = limited.Write(data)

	if time.Now().Sub(start) <= time.Second {
		t.Error("The limiter didn't block when it should have")
	}
}

func TestRateLimitReadWriteCloserReadAndWrite(t *testing.T) {
	data := make([]byte, 1024)
	limit := 1024
	limited := NewRateLimitReaderWriterCloser(&dummyReadWriter{}, uint64(limit))

	start := time.Now()

	for i := 0; i < limit; i++ {
		var err error
		if i%2 == 0 {
			_, err = limited.Read(data)
		} else {
			_, err = limited.Write(data)
		}
		assert.NoError(t, err)
	}

	if time.Now().Sub(start) >= time.Second {
		t.Error("The limiter blocked when it shouldn't have")
	}
}
