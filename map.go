package main

import "fmt"

func main() {
	map1()
	map2()
}

func map1() {

	var map1 map[string]int = map[string]int{"english": 90, "math": 80}
	fmt.Println(map1)

	map2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(map2)

	score := make(map[string]int)
	score["english"] = 90
	score["math"] = 80
	fmt.Println(score)
	fmt.Printf("score的长度为: %d\n, 类型%T\n", len(score), score)

	var map3 map[string]int

	map3 = map[string]int{"a": 1, "b": 2}

	fmt.Println(map3)
}

func map2() {
	map2 := map[string]int{"name": 1, "age": 18}
	map2["name1"] = 2
	map2["age1"] = 19
	fmt.Println(map2)

	delete(map2, "name")
	fmt.Println(map2)
	fmt.Printf("map2的长度为: %d\n, %d", len(map2), map2["s1"])
	fmt.Println("-----------------------")
	fmt.Println(map2["s1"])
	s2, ok := map2["s1"]
	fmt.Println(s2, ok)
	if ok {
		fmt.Println("存在", map2["s1"])
	} else {
		fmt.Println("不存在")
	}

	if s1, ok := map2["s1"]; ok {
		fmt.Println("存在", s1)
	} else {
		fmt.Println("不存在")
	}
}
