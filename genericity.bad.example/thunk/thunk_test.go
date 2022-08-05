package thunk

import "testing"

func TestNewThunk(t *testing.T) {
	cache := make([]*Thunk[int], 41)
	var fib func(int) int
	fib = func(n int) int {
		return cache[n-1].Force() + cache[n-2].Force()
	}
	for i := range cache {
		i := i
		cache[i] = NewThunk(func() int {return fib(i)})
	}
	cache[0].o.Set(0)
	cache[1].o.Set(1)
	t.Log(cache[40].Force())
}