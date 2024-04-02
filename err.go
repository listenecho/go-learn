package main

import "fmt"

func main() {
	fmt.Println("c")
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		// if err := recover(); err != nil {
		// 	fmt.Println(err, 44) // 这里的err其实就是panic传入的内容，55
		// }
		fmt.Println("e333")
	}()

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		// if err := recover(); err != nil {
		// 	fmt.Println(err, 333) // 这里的err其实就是panic传入的内容，55
		// }
		fmt.Println("e2222")
	}()

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		// if err := recover(); err != nil {
		// 	fmt.Println(err, 5555) // 这里的err其实就是panic传入的内容，55
		// }
		fmt.Println("e111")
	}()
	f()              //开始调用f
	fmt.Println("f") //这里开始下面代码不会再执行
}

func f() {
	fmt.Println("a")
	panic("this is 1")
	panic("this is 2")
	panic("this is 3")

	fmt.Println("b") //这里开始下面代码不会再执行
	fmt.Println("f")
}

// painc 用于处理异常错误
// go 语言没有 try catch finally 设计者认为异常与控制结构混在一起会使代码边的异常
// Exception处理：defer, panic, recover。

// Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
