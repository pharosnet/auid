package auid

import (
	"sync"
	"math"
	"errors"
	"sync/atomic"
)

const atomic_num_add_retry_times = 64

// AtomicNumber new one with default
func NewAtomicNumber() *AtomicNumber {
	n := new(AtomicNumber)
	n.number = uint64(0)
	n.retry = atomic_num_add_retry_times
	n.mu = new(sync.Mutex)
	return n
}

// AtomicNumber number(uint64) increase by cas and mutex when over retry times
type AtomicNumber struct {
	number uint64
	retry int
	mu *sync.Mutex
}

// increase number by cas and mutex when over retry times
func (n *AtomicNumber) Increase() (uint64, error) {
	retry := atomic_num_add_retry_times
	if n.number == math.MaxUint64 {
		return n.number, errors.New("number is max.")
	}
	for !atomic.CompareAndSwapUint64(&n.number, n.number, n.number + uint64(1)) {
		retry --
		if retry > 0 {
			continue
		}
		n.mu.Lock()
		n.number = n.number + uint64(1)
		n.mu.Unlock()
		break
	}
	return n.number, nil
}

// set 0 into number
func (n *AtomicNumber) Reset() {
	n.mu.Lock()
	n.number = uint64(0)
	n.mu.Unlock()
}

