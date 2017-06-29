package auid

import (
	"sync"
	"errors"
	"sync/atomic"
)

const atomic_num_add_retry_times = 64
const atomic_num_max_value = int64(99999999999)
const atomic_num_min_value = int64(10000000000)


// AtomicNumber new one with default
func NewAtomicNumber() *AtomicNumber {
	n := new(AtomicNumber)
	n.number = atomic_num_min_value
	n.retry = atomic_num_add_retry_times
	n.mu = new(sync.Mutex)
	return n
}

// AtomicNumber number(int64) increase by cas and mutex when over retry times
type AtomicNumber struct {
	number int64
	retry int
	mu *sync.Mutex
}

// increase number by cas and mutex when over retry times
func (n *AtomicNumber) Increase() (int64, error) {
	retry := atomic_num_add_retry_times
	if n.number > atomic_num_max_value {
		return n.number, errors.New("number is max.")
	}
	for !atomic.CompareAndSwapInt64(&n.number, n.number, n.number + int64(1)) {
		retry --
		if retry > 0 {
			continue
		}
		n.mu.Lock()
		n.number = n.number + int64(1)
		n.mu.Unlock()
		break
	}
	return n.number, nil
}

// set 0 into number
func (n *AtomicNumber) Reset() {
	n.mu.Lock()
	n.number = atomic_num_min_value
	n.mu.Unlock()
}

