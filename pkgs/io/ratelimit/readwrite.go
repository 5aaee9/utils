package ratelimit

import (
	"io"
)

type RateLimitReadWriteCloser struct {
	limiter  *Limiter
	underlay io.ReadWriteCloser
}

func NewRateLimitReaderWriterCloser(underlay io.ReadWriteCloser, size uint64) io.ReadWriteCloser {
	return &RateLimitReadWriteCloser{
		limiter:  NewLimiter(size),
		underlay: underlay,
	}
}

func NewRateLimitReaderWriterCloserShared(underlay io.ReadWriteCloser, l *Limiter) io.ReadWriteCloser {
	return &RateLimitReadWriteCloser{
		limiter:  l,
		underlay: underlay,
	}
}

func (r *RateLimitReadWriteCloser) Close() error {
	return r.underlay.Close()
}

func (r *RateLimitReadWriteCloser) Read(p []byte) (int, error) {
	n, err := r.underlay.Read(p)
	if err != nil {
		return n, err
	}

	r.limiter.Wait(uint64(n))

	return n, nil
}

func (r *RateLimitReadWriteCloser) Write(p []byte) (int, error) {
	n, err := r.underlay.Write(p)
	if err != nil {
		return n, err
	}

	r.limiter.Wait(uint64(n))

	return n, nil
}
