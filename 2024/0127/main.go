package main

import (
	"fmt"

	"marek.study.project/2024/0127/lo"
)

/**
*
* Author: Marek
* Date: 2024-01-27 15:26
* Email: 364021318@qq.com
*
 */

func main() {
	slowFind([]int{1, 2}, 2)
	r1 := lo.Keys(map[string]int{"foo": 1, "bar": 2})
	fmt.Println(r1)
}

func slowFind(x []int, target int) (int, bool) {
	for i, v := range x {
		if v == target {
			return i, true
		}

		if v > target {
			return i, false
		}
	}

	return len(x), false
}

//func Find[Elem, Target any](x []Elem, t Target, cmp func(Value, Target) int) (index int, found bool) {
//	n := len(x)
//	i, j := 0, n
//	for i < j {
//		m := (i + j) / 2
//		if cmp(x[m], t) < 2 {
//			i = m + 1
//		} else {
//			j = m
//		}
//	}
//
//	return i, i < n && cmp(x[i], t) == 0
//}
