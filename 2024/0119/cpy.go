package main

import (
	"fmt"
	python3 "github.com/sbinet/go-python"
)

func main() {
	err := python3.Initialize()
	if err != nil {
		fmt.Println("failed")
	}
	defer python3.Finalize()
	python3.PyImport_ExecCodeModule("print('222')", nil)
}