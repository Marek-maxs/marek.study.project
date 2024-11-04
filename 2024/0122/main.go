package main

import "fmt"

/**
*
* Author: Marek
* Date: 2024-01-22 20:20
* Email: 364021318@qq.com
*
 */

func main() {
	nodeMap := make(map[string]string)

	nodeMap["1"] = "a"
	nodeMap["2"] = "b"

	for key := range nodeMap {
		fmt.Println(key)
	}
}