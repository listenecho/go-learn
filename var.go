package main

import "fmt"

// 全局变量

var global1 = 100
var global2 = 1001

// 一次性声明全局变量
var (
	global3 = 1003
	global4 = 999
)

func main() {
	var age int
	age = 111

	fmt.Println("age", age)

	// 变量类型必须和值的类型相互匹配
	var num float32
	num = 0.1
	fmt.Println("age", num)

	// 指定的变量类型,不赋值,将使用默认值
	var age1 int
	fmt.Println(age1) // 0

	// 定义变量如果没有显示指定变量类型, 则会根据后面的值进行类型推断
	var age2 = 1
	fmt.Println(age2)
	// 变量在声明的时候必须指定变量类型或者给一个初始值
	// var age3
	// age3 = 1
	// fmt.Println(age3)

	// 省略var 关键字声明变量,也是自动推荐值的类型
	sex := "男"
	fmt.Println("性别", sex)

	fmt.Println("-----------------------------------------")

	// 声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n11, n22, n33 = "10", 1, 10.2
	fmt.Println(n11)
	fmt.Println(n22)
	fmt.Println(n33)

	n111, n232, n334 := "10", 1, 10.2
	fmt.Println("------------")
	fmt.Println(n111)
	fmt.Println(n232)
	fmt.Println(n334)
	fmt.Println("------------")

	// 使用全局变量
	fmt.Println(global1)
	fmt.Println(global2)

	fmt.Println(global3)
	fmt.Println(global4)
	// 声明多个变量
	var (
		global5 = 1003111
		global6 = "999"
	)

	fmt.Println(global5)
	fmt.Println(global6)

	fmt.Println("-----指针-------")
	var ptr = &global1
	fmt.Println("指针", ptr)

	ptr2 := new(int)
	*ptr2 = 200
	fmt.Println("指针2", ptr2)
	fmt.Println("指针2", *ptr2)

}
