package main

import "fmt"

// Its zore value is not a valid ID.
type ID uint64

type IDGenerator struct {
	IDFn func() ID
}

func (g IDGenerator) ID() ID {
	return g.IDFn()
}

func NewStaticIDGenerator(id ID) IDGenerator {
	return IDGenerator{
		IDFn: func() ID { return id },
	}
}

func main() {
	data := NewStaticIDGenerator(ID(100))

	fmt.Println(data.ID())
}
