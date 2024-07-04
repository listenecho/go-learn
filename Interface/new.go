package main

import (
	"fmt"
)

type UI string

func (d UI) getName() {
	fmt.Println(d)
}

func main() {
	var u UI
	u = UI("1111")
	fmt.Println(u)
}
