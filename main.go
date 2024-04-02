package main

import (
	"fmt"
)

// 变量名 变量类型 没有初始值的情况下 int的默认值为0 string类型的默认值为 ""
func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// 一次声明多个变量
func variableMore() {
	var a, b int = 4, 5
	fmt.Printf("%d %d\n", a, b)
}

// 省略类型声明 会根据初始值推断出
// 省略var，使用 :=

func variableGuess() {
	var a = 1
	b := "username" // 这样声明变量的方式只是在函数内部实现，一遍全局变量不能使用:=声明
	fmt.Printf("%d %s", a, b)
}

// 定义常量
func variableConst() {
	const username = "wenbobno"
	fmt.Printf("\n %s", username)
}

// 使用常量定义枚举
// iota 表示自增加 _跳过

func enums() {
	const (
		cpp = iota
		java
		python
		javascript
	)
	fmt.Println(cpp, java, python, javascript)
}

func enums2() {
	const (
		cpp = iota
		_
		java
		_
		python
	)
	fmt.Println("\n", cpp, java, python)
}

// 类型转换 go只有强制类型转换 没有隐式类型转换
func tranfrom() {
	a, b := 4, 9
	var c int
	c = int(float64(1*4 + a*a + b*b))
	fmt.Println(c, "this is c")
}

func main() {
	fmt.Println("yyyy")
	variable()
	variableMore()
	variableGuess()
	variableConst()
	enums()
	enums2()
	tranfrom()
}
