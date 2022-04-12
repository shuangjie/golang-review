package main

import "fmt"

//寻找最长不含有重复字符的子串
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	fmt.Println("lastOccurred = ", lastOccurred)
	for i, ch := range []rune(s) {
		fmt.Println("i = ", i)
		fmt.Println("ch = ", ch)
		fmt.Println("start = ", start)
		fmt.Println("maxLength = ", maxLength)
		fmt.Println("lastOccurred[ch] = ", lastOccurred[ch])
		LastO, okk := lastOccurred[ch]
		fmt.Println("LastO = ", LastO)
		fmt.Println("okk = ", okk)

		if LastI, ok := lastOccurred[ch]; ok && LastI >= start {

			start = LastI + 1
			fmt.Println("LastI = ", LastI)
			fmt.Println("start = ", start)
		}
		if i-start+1 > maxLength {
			fmt.Println("i-start+1 > maxLength")
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
		fmt.Println(i, ch)
		fmt.Println("lastOccurred = ", lastOccurred)
		fmt.Println()

	}
	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	//fmt.Println(
	//	lengthOfNonRepeatingSubStr("abcdef"))
	//fmt.Println(
	//	lengthOfNonRepeatingSubStr("b"))
	//fmt.Println(
	//	lengthOfNonRepeatingSubStr("广东省123123"))
}
