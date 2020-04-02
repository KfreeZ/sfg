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

		arr = []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
		getLeastNumbers(arr, 8)
	}
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	quickSelect(arr, 0, len(arr)-1, k-1)

	return arr[:k]
}

func quickSelect(arr []int, lo int, hi int, k int) {

	pivotIndex := partition(arr, lo, hi)

	if pivotIndex == k {
		return
	}

	fmt.Println("pivot: ", pivotIndex)
	if pivotIndex > k {
		quickSelect(arr, lo, pivotIndex-1, k)
	} else {
		quickSelect(arr, pivotIndex+1, hi, k)
	}
}

func partition(arr []int, lo int, hi int) int {
	i := lo
	v := arr[lo]
	j := hi
	for true {
		for ; i < hi; i++ {
			if arr[i] > v {
				break
			}
		}
		for ; j > lo; j-- {
			if arr[j] < v {
				break
			}
		}
		if i >= j {
			break
		}
		swap(arr, i, j)
	}
	swap(arr, lo, j)

	return j
}

func swap(arr []int, i int, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t

	fmt.Println(stringutils.IsValid("["))

}
