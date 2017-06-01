package auid

import (
	"sync"
	"math"
	"errors"
	"sync/atomic"
)

const atomic_num_add_retry_times = 64

func NewAtomicNumber() *AtomicNumber {
	n := new(AtomicNumber)
	n.val = uint64(0)
	n.retry = atomic_num_add_retry_times
	n.mu = new(sync.Mutex)
	return n
}

type AtomicNumber struct {
	val uint64
	retry int
	mu *sync.Mutex
}

func (n *AtomicNumber) Add() (uint64, error) {
	retry := atomic_num_add_retry_times
	if n.val == math.MaxUint64 {
		return n.val, errors.New("number is max.")
	}
	for !atomic.CompareAndSwapUint64(&n.val, n.val, n.val + uint64(1)) {
		retry --
		if retry > 0 {
			continue
		}
		n.mu.Lock()
		n.val = n.val + uint64(1)
		n.mu.Unlock()
		break
	}
	return n.val, nil
}

func (n *AtomicNumber) Reset() {
	n.mu.Lock()
	n.val = uint64(0)
	n.mu.Unlock()
}

