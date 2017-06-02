package auid

import (
	"sync"
)

var pool sync.Pool
var once sync.Once

// new auid from pool
func NewAuidWithPool() string {
	id := get()
	val := id.next()
	put(id)
	return val
}

// get auid from pool
func get() *auid {
	once.Do(func() {
		pool = sync.Pool{New:func() interface{} {
			return &auid{
				uuid:NewUUIDV4(),
				number:NewAtomicNumber(),
			}
		}}
	})
	id, ok := pool.Get().(*auid)
	if !ok || id == nil {
		return &auid{
			uuid:NewUUIDV4(),
			number:NewAtomicNumber(),
		}
	}
	return id
}

// set auid into pool
func put(id *auid)  {
	pool.Put(id)
}