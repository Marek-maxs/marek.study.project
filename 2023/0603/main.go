package main

import "sync"

type TestStruct struct {
	FiledD interface{}
	*int
	Date   string
	DD     string
	Slices []string
	Ints   [3]int
	FieldA int64
	FieldB float64
	sync.Mutex
	FieldC float32
}

type OldTestStruct struct {

	FieldC float32 // align:4 size:4

	FieldB float64 // align:4 size:8

	sync.Mutex // align:4 size:8

	FieldA int64 // align:8 size:8

	Ints [3]int // align:wordsize size:24

	*int // align:8 haspointer size:8

	Date string // align:2*wordsize haspointer size:2*wordsize

	FiledD interface{} // align:2*wordsize haspointer size:wordsize*2

	Slices []string // align:3*wordsize haspointer size:3*wordsize

}
