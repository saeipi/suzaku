package main

import "fmt"

func main() {
	test0()
	fmt.Println("代码来到了这里")
	var input int
	fmt.Scan(&input)
}

func test0() {
	defer func() {
		if err := recover(); &err != nil {
			fmt.Println(err)
		}
	}()
	test1()
}

func test1() int {
	var list = make([]int, 0)
	var val = list[1]
	return val
}
