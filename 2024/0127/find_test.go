package main

import "testing"

/**
*
* Author: Marek
* Date: 2024-01-27 15:38
* Email: 364021318@qq.com
*
 */

var findTests = []struct{x []int; t int; i int; ok bool} {
	{[]int{}, 100, 0, false},
	{[]int{10, 20, 30}, 10, 0, true},
	{[]int{10, 20, 30}, 30, 2, true},
	{[]int{10,20,30}, -10, 0, false},
	{[]int{10, 20, 30}, 50, 3, false},
	{[]int{10, 20, 30}, 25, 2, false},
}

func TestSlowFind(t *testing.T) {
	for _, e := range findTests {
		i, ok := slowFind(e.x, e.t)
		if i != e.i || ok != e.ok {
			t.Errorf("slowFind(%v, %v) = %v, %v, want %v, %v", e.x, e.t, i, ok, e.i, e.ok)
		}
	}
}