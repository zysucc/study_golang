package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	s := arr[2:6]
	d := arr[:]
	fmt.Println(s, d)
	s1 := arr[2:]
	s2 := arr[:]
	//	updateSlice(s1)
	fmt.Println(s1)
	//	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println("Reslice")
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
	fmt.Println(cap(s1))

	fmt.Println("new string")

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println(s3, s4, s5)
	fmt.Println(arr)

	fmt.Println(cap(arr))
}
