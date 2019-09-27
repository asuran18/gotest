package main

import (
	"fmt"
)

/**
 * 值交换
 */
func main() {
	var a int = 100
	var b int = 200
	b, a = a, b
	fmt.Println(a, b)
}
