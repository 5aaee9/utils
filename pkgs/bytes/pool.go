package bytes

import "sync"

type BytesPool struct {
	size uint
	pool sync.Pool
}

func NewBytesPool(size uint) *BytesPool {
	return &BytesPool{
		size: size,
		pool: sync.Pool{
			New: func() any {
				return make([]byte, size)
			},
		},
	}
}

func (p *BytesPool) Get() []byte {
	return p.pool.Get().([]byte)
}

func (p *BytesPool) Put(data []byte) {
	if uint(len(data)) != p.size {
		panic("buf size not match pool size")
	}

	p.pool.Put(data)
}
