package main

import (
	"fmt"
)

type UI interface {
	getName() string
}

type User struct {
	name string
	age  int
}

func (u User) getName() string {
	return u.name
}

func main() {
	fmt.Println("hello world")
	var u UI

	u = User{
		name: "wenbono",
		age:  18,
	}

	text := u.getName()

	fmt.Println(text)

	excuse()
}

type Duck interface {
	getDuckName()
}

type DuckName string

func (d DuckName) getDuckName() {
	fmt.Println(d)
}

type RuberDuck string

func (d RuberDuck) getDuckName() {
	fmt.Println(d)
}

type Swan struct {
	name string
}

func (s Swan) getDuckName() {
	fmt.Println(s.name)
}

func excuse() {
	var d Duck
	d = DuckName("duck")
	d.getDuckName()

	d = RuberDuck("ruber duck")
	d.getDuckName()

	d = Swan{
		name: "swan",
	}
	d.getDuckName()
}
