package main

import (
	"fmt"
	"runtime"
)

/**
 * go测试类
 */
func main() {
	println("Hello", "world")
	fmt.Printf("%s", runtime.Version())
}
