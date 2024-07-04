package main

import "fmt"

func main() {
	array1()
	array2()
}

func array1() {
	var a [10]int

	fmt.Println(a)

}

// 需要注意: 由于数组长度是固定的，因此在 Golang 中很少直接使用数组。更常见的是使用切片（slice），它是一个动态数组，可以根据需要动态增加或减少大小。
func array2() {

	var arr [3]int = [3]int{1, 2, 3}

	arr2 := [4]int{1, 2}

	arr3 := [...]int{1, 2, 3, 4, 5}

	arr4 := [3]int{1: 1, 2: 4}

	type arrType [3]int

	arr5 := arrType{1, 2, 3}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	arr3[0] = 100
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

	arr6 := [3]bool{true, false}
	fmt.Println(arr6)

}
