package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StructDemo1 struct {
	name string
	age  int
	B    byte
	book []string
}

func main() {
	Demo1()
	Demon()
	Demo2()
	Demo4()
	Demo5()
}

func Demo1() {
	me1 := StructDemo1{
		name: "zhangsan",
		age:  18,
		B:    'B',
		book: []string{"book1", "book2"},
	}
	fmt.Printf("\td value %+v\n", me1)

	me1.age = 333
	fmt.Printf("\tage value %+v\n", me1)

	fmt.Printf("\tage value %+v\n", me1.age)

}

// 结构体Tag
// 结构体Tag是结构体字段的元信息，可以在运行的时候通过反射的机制读取出来。
// 结构体Tag的格式是`key:"value"`，key是一个不带引号的字符串，value可以是带引号的字符串。
// 结构体Tag是结构体字段的一部分，因此必须放在字段的后面。
// 结构体Tag的解析代码在reflect包中的StructTag类型中。
// 结构体Tag的一个重要应用场景是序列化和反序列化。

func Demon() {
	type USer struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := USer{
		Name: "zhangsan",
		Age:  18,
	}
	fmt.Printf("\tu value %+v\n", user)

	bytes, err := json.Marshal(user)

	if err != nil {
		fmt.Println("json marshal failed")
		return
	}

	fmt.Printf("\tjson value %+v\n", string(bytes))

}

func Demo2() {
	type User struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
	}

	user := User{
		UserName: "zhangsan",
		Age:      18,
	}

	t := reflect.TypeOf(user)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name: %s, type: %s, json tag: %s\n", field.Name, field.Type, field.Tag.Get("json"))
	}
}

type User struct {
	UserName string
	Age      int
}

func Demo4() {

	data := User{
		UserName: "zhangsan",
		Age:      18,
	}
	fmt.Println(data)

	data.getB()
	data.getAA(188)
	getBBB(data, 98)
	fmt.Println(data)

}

func (d User) getB() {
	d.Age = 100
	fmt.Printf("d.Age %+v\n", d.Age)
}

func (d User) getAA(age int) {
	d.Age = age
	d.UserName = "wenbono"
	fmt.Printf("d.Age %+v\n", d.Age)
}

func getBBB(data User, age int) {
	data.Age = age
	data.UserName = "wenbono"
	fmt.Printf("d.Age %+v\n", data.Age)
}

func Demo5() {
	u := User{
		UserName: "zhangsan",
		Age:      18,
	}
	u.ptrFun()
	fmt.Printf("d.Age %+v\n", u.Age)
}

func (u *User) ptrFun() {
	u.Age = 100
	fmt.Printf("u.Age %+v\n", u.Age)
}
