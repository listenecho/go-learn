package main

import "fmt"

func main() {

	// grade(55)
	score(100)
}

func grade(score int) string {
	g := ""
	defer func() {

		switch {
		case score > 10:
			g = "2222"
			// panic(fmt.Sprintln("bad"))
		case score <= 10:
			g = "jaja"
		}

	}()
	f()
	return g

}

func f() {
	fmt.Println("a")
	panic("异常信息")
	//这里开始下面代码不会再执行

}

// painc 用于处理异常错误
// go 语言没有 try catch finally 设计者认为异常与控制结构混在一起会使代码边的异常
// Exception处理：defer, panic, recover。

func score(score int) {
	switch score / 10 {
	case 9:
		fmt.Println("A")
	default:
		fmt.Println("无数据")
	}
}
