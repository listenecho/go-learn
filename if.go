package main

import (
	"fmt"
	"io/ioutil"
)

// fmt.Print === 原样输出
// fmt.Printf === 格式输出
// fmt.Println === 值 + 空格 输出
func main() {
	fmt.Println("this is IF demo")
	if1()
	if2()
	if3()
}

func if1() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func if2() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contents)
	}
}

func if3() {
	if count := 10; count > 100 {
		fmt.Println("count 大于 100")
	} else {
		fmt.Println("count 小于 100")
	}
}
