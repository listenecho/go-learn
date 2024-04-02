package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var number1 int8 = 100
	fmt.Println("number1", number1)

	// var number2 int8 = 199
	// fmt.Println("number1", number2)

	var number3 = 1000000000000000000
	fmt.Printf("number1 %T", number3)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(number3))
}

/**
数据类型

  基本数据类型
			数值型
					整数
					浮点
			字符型
			布尔型
			字符串型

	派生数据类型 / 复杂数据类型
		 指针
		 数组
		 结构体
		 管道 ?
		 函数 ?
		 切片 ?
		 接口 ?
		 map ?
*/
