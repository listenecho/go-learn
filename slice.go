package main

import "fmt"

func main() {
	slice1()

	slice2()

	slice3()

	slice4()

	arr1()
}

func arr1() {
	arr := [...]int{1, 1, 1}
	fmt.Printf("数组长度 %d, 容量 %d\n， 类型 %T\n", len(arr), cap(arr), arr)
}

func slice1() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("数组长度 %d, 容量 %d\n", len(arr), cap(arr))

	slice := arr[1:3]
	fmt.Printf("切片长度 %d, 容量 %d\n", len(slice), cap(slice))

	slice = arr[1:3:4]
	fmt.Printf("切片长度 %d, 容量 %d\n", len(slice), cap(slice))

	slice = arr[1:3:5]
	fmt.Printf("切片长度 %d, 容量 %d\n", len(slice), cap(slice))
}

func slice2() {
	// 一个切片具备的三个要素：类型（Type），长度（size），容量（cap）

	var sl1 []int
	var sl2 []string
	var sl3 = []int{}

	fmt.Println(sl1, sl2, sl3)
}

func slice3() {
	slice := make([]int, 10, 19)
	fmt.Printf("切片长度 %d, 容量 %d\n", len(slice), cap(slice))

}

func slice4() {
	fmt.Println("-----------------")
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("numbers4的长度为: %d\n", len(numbers4))
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d\n, 容量为 %d\n", myslice, len(myslice), cap(myslice))

	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}
