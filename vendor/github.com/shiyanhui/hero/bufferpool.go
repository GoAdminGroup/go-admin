package hero

import (
	"bytes"
	"sync"
)

const buffSize = 10000

var defaultPool *pool

func init() {
	defaultPool = newPool()
}

type pool struct {
	pool *sync.Pool
	ch   chan *bytes.Buffer
}

func newPool() *pool {
	p := &pool{
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
		ch: make(chan *bytes.Buffer, buffSize),
	}

	// It's faster with unused channel buffer in go1.7.
	// TODO: need removed?
	for i := 0; i < buffSize; i++ {
		p.ch <- new(bytes.Buffer)
	}

	return p
}

// GetBuffer returns a *bytes.Buffer from sync.Pool.
func GetBuffer() *bytes.Buffer {
	return defaultPool.pool.Get().(*bytes.Buffer)
}

// PutBuffer puts a *bytes.Buffer to the sync.Pool.
func PutBuffer(buffer *bytes.Buffer) {
	if buffer == nil {
		return
	}

	buffer.Reset()
	defaultPool.pool.Put(buffer)
}
