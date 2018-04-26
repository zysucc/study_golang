package main

import (
	"fmt"
	"math"
	"os"
)

func consts() {
	const filename = "123.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		java
		php
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(cpp, java, php)
	fmt.Println(b, kb, mb, gb, tb)
}

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	consts()
	enums()
}
