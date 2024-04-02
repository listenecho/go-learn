package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int
	size := unsafe.Sizeof(x)
	fmt.Printf("int 类型的大小为 %d 字节\n", size)
}
