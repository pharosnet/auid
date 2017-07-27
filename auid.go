package auid

import (
	"fmt"
)

func init()  {
	id = &auid{
		uuid:NewUUIDV4(),
		number:NewAtomicNumber(),
	}
}

var id *auid

// use global auid to gen a id.
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx-xxxxxxxxxxx.
// len = 48 
func NewAuid() string {
	val := id.next()
	return val
}

// AUID base on UUIDv4 and AtomicNumber
type auid struct {
	uuid UUID
	number *AtomicNumber
}

// AUID call AtomicNumber.Increase() to get next number
func (a *auid) next() string {
	n, err := a.number.Increase()
	if err != nil {
		a.uuid = NewUUIDV4()
		a.number.Reset()
		return a.next()
	}
	return fmt.Sprintf("%s-%d", a.uuid.String(), n)
}
