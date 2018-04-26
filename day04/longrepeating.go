package main

import "fmt"

func tests(s string) int {
	lastOccurred := make(map[byte]int)
	fmt.Println(lastOccurred)
	fmt.Println([]byte(s))
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		fmt.Println(lastI)
		fmt.Println(ok)
		if ok && lastI >= start {
			start = lastI + 1
		}
		fmt.Println("-------------------")
		fmt.Println(start)
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
		fmt.Println("-------------------")
		fmt.Println(lastOccurred)
	}
	return maxLength
}

func main() {
	fmt.Println(tests("abcdeabcghijk"))

}
