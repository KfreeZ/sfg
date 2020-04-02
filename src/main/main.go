package main

import (
	"fmt"
	"heap"
	"stack"
	"stringutils"
)

func main() {

	fmt.Println("Go strive for greatness")

	as := stack.NewArrayStack()
	as.IsEmpty()

	if false {

		rets := stringutils.CompressString("aabcccccaa")
		fmt.Print(rets)

		fmt.Print(stringutils.LengthOfLongestSubstring("pwwkew"))

		words := []string{"hello", "world", "leetcode"}
		chars := "welldonehoneyr"
		fmt.Print(stringutils.CountCharacters(words, chars))

		stringutils.MyAtoi(" ")

		arr := []int{3, 2, 1}
		fmt.Println(heap.GetLeastNumbers(arr, 2))
		arr = []int{0, 1, 2, 1}
		fmt.Println(heap.GetLeastNumbers(arr, 1))
	}

	fmt.Println(stringutils.IsValid("["))

}
