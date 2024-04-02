package main

import (
	"fmt"
	"go-learn/util2"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "/":
		return a / b, nil
	case "*":
		return a * b, nil
	default:
		return 0, fmt.Errorf("op: %s", op)
	}

}

// 指定函数的返回值
func div2(a, b int) (q, r int) {
	q = a / b
	r = b * a
	return
}

func calc(a int, b int) int {
	return a + b
}

func main() {
	util2.Add(1, 4)
	fmt.Println(eval(16, 4, "/"))
	fmt.Println(div2(10, 5))
	sum := calc(100, 90)
	fmt.Println(sum)
}
