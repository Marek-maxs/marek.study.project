package main

import (
	"reflect"
	"testing"
)

/**
*
* Author: Marek
* Date: 2023-04-15 11:27
* Email: 364021318@qq.com
*
 */

func TestNewHello(t *testing.T) {
	tests := []struct {
		name string
		want *Hello
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHello(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHello() = %v, want %v", got, tt.want)
			}
		})
	}
}