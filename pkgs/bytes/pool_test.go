package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoolGet(t *testing.T) {
	pool := NewBytesPool(200)
	assert.Equal(t, 200, len(pool.Get()))

	pool = NewBytesPool(1025)
	assert.Equal(t, 1025, len(pool.Get()))

	pool = NewBytesPool(2 * 1024)
	assert.Equal(t, 2*1024, len(pool.Get()))

	pool = NewBytesPool(5 * 2000)
	assert.Equal(t, 5*2000, len(pool.Get()))
}

func TestPoolPut(t *testing.T) {
	pool := NewBytesPool(200)

	assert.Panics(t, func() {
		pool.Put(make([]byte, 2000))
	})

	assert.NotPanics(t, func() {
		pool.Put(make([]byte, 200))
	})
}
