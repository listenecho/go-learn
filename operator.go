package main

import "fmt"

func main() {
	ArithmeticOperator()
	assignmentOperator()
	pointerOperator()
}

// 算数运算符
func ArithmeticOperator() {
	// +
	// 1. 正数 2. 加法运算 3. 字符串拼接
	fmt.Println("+ 加号")
	var n1 int = +10
	fmt.Println(n1)

	var n2 int = 3
	fmt.Println(n2 + n1)

	var n3 string = "10" + "xxxxxx"
	fmt.Println(n3)

	// /除号
	fmt.Println(" / 除号")
	var n4 = 10 / 3
	fmt.Printf("int类型 相除的到的数据类型：%T \n", n4)
	fmt.Println(10 / 3) // 两个Int类型数据相除得到的是整数类型
	var n5 = 10.0 / 3
	fmt.Printf("浮点类型参与计算得到的数据类型：%T \n", n5)
	fmt.Println(10.0 / 3) // 浮点类型的数据参与计算，得到的值为浮点类型

	// % 取模操作 等价公式 a%b = a - a/b*b
	fmt.Println("% 取模 ")
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)

	// ++ 自增 -- 自减
	fmt.Println("++ 自增 -- 自减 ")
	var n6 int = 10
	n6++
	fmt.Println(n6)
	//go语言中 ++ -- 只能单独使用， 不能参与到运算中
	//go语言中 ++ -- 只能使用在变量的后面， 不能再变量的前面使用， 错误用法：++a --a
}

// 赋值运算符
// 将某个运算后的值，赋值给指定的变量
func assignmentOperator() {
	// = 等号 = 右侧得到计算结果之后， 再赋值给左侧的变量

	// +=    a = a + 10 ==> a += 10
	// /=    a = a / 10 ==> a /= 10
	// %=    a = a % 10 ==> a %= 10

	// -=*=
}

// 关系运算符

// 逻辑运算符

// 位运算符

// 其他元素符

func pointerOperator() {
	var name int = 10
	fmt.Println("name 的指针地址为: ", &name)

	var ptr *int = &name
	fmt.Println("ptr 存储的指针地址为: ", ptr)
	fmt.Println("prt 这个指针变量存储的值为", *ptr)
}
