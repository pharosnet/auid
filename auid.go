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

func NewAuid() string {
	val := id.next()
	return val
}

type auid struct {
	uuid UUID
	number *AtomicNumber
}

func (a *auid) next() string {
	n, err := a.number.Add()
	if err != nil {
		a.uuid = NewUUIDV4()
		a.number.Reset()
		return a.next()
	}
	return fmt.Sprintf("%s:%d", a.uuid.String(), n)
}
