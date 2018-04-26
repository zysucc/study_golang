package main

import "fmt"

func test() {
	var a int
	var s string
	fmt.Println(a, s)
}

func test1() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func test2() {
	a, b, c, d := 3, 4, true, "daskd"
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("vim-go")
	test()
	test1()
	test2()
}
