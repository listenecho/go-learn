package main

import (
	"fmt"
	"time"
)

func main() {

	language := []string{"golang", "java", "c++", "python", "rust", "js"}
	level := []string{"junior", "middle", "senior"}

	go showLanguage(language)
	go showLevel(level)

	<-time.After(time.Second * 1)
	fmt.Println("return")

}

func showLanguage(language []string) {
	for _, v := range language {
		fmt.Println(v)
		time.Sleep(time.Second)

	}
}

func showLevel(level []string) {
	for _, v := range level {
		fmt.Println(v)
		time.Sleep(time.Second)

	}
}
