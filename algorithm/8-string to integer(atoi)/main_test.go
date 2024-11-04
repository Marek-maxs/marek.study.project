package main

import "testing"

func TestMyAtoi(t *testing.T) {
	got := MyAtoi("string")
	want := 0

	if got != want {
		t.Error("MyAtoi want")
	}
}