package thunk

import (
	"marek.study.project/genericity.bad.example/option"
	)

type Thunk[T any] struct {
	doer func() T // action being thunked
	o option.Option[T]  // cache for complete thunk data
}

func (t *Thunk[T]) Force() T {
	if t.o.IsSome() {
		return t.o.Yank()
	}

	t.o.Set(t.doer())
	return t.o.Yank()
}

func NewThunk[T any] (doer func() T) *Thunk[T] {
	return &Thunk[T] {
		doer: doer,
		o: option.NewOption[T](),
	}
}