package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len = %d,cap = %d\n", len(s), cap(s))
}

func main() {
	var s []int //zero value for sliceis nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	fmt.Println(s2)
	fmt.Println("new slice copy clice")
	copy(s2, s1)
	fmt.Println(s2)
	s3 := make([]int, 10, 32)

	printSlice(s2)
	printSlice(s3)
}
