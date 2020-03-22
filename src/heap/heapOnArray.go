package heap

import (
	"fmt"
)

func GetLeastNumbers(arr []int, k int) []int {
	maxTopHeap := NewMaxHeapOnArray(k)
	fmt.Println(maxTopHeap.Data)

	for _, v := range arr {
		if !maxTopHeap.IsFull() {
			maxTopHeap.Push(v)
		} else if maxTopHeap.Top() > v {
			maxTopHeap.Pop()
			maxTopHeap.Push(v)
		}
	}

	return maxTopHeap.Data[1:]
}

// turn the heap(min/max) the childrens' is father's index*2 and index*2+1
// but the index 0 should be reserved

type heapOnArray struct {
	Data     []int
	Cnt      int
	capacity int
}

func NewMaxHeapOnArray(capa int) *heapOnArray {
	return &heapOnArray{
		capacity: capa,
		Data:     make([]int, capa+1),
		Cnt:      0,
	}
}

// Insert into the MAX top heap,
// if the heap is full
// if the the new value is bigger than the tail value,
// it would replace one value in the heap
// or the value would be dropped if it is the smallest
func (ha *heapOnArray) Push(newItem int) bool {
	fmt.Println(ha.Data, "| push ", newItem, "| Cnt ", ha.Cnt)

	if ha.Cnt == ha.capacity+1 {
		return false
	}

	ha.Cnt++
	newIndex := ha.Cnt
	ha.Data[newIndex] = newItem
	fatherIndex := newIndex / 2

	for ha.Data[newIndex] > ha.Data[fatherIndex] && fatherIndex > 0 {
		swap(ha.Data, newIndex, fatherIndex)
		newIndex = fatherIndex
		fatherIndex = newIndex / 2
	}

	return true
}

func (ha *heapOnArray) Top() int {
	return ha.Data[1]
}

func (ha *heapOnArray) IsFull() bool {
	return ha.Cnt == ha.capacity
}

func (ha *heapOnArray) Pop() {

	if ha.Cnt == 0 {
		return
	}

	fmt.Println(ha.Data, "| pop  ", ha.Data[ha.Cnt], "| Cnt ", ha.Cnt)

	swap(ha.Data, 1, ha.Cnt)
	ha.Data[ha.Cnt] = 0
	ha.Cnt--

	ha.Heapfy()
}

func (ha *heapOnArray) Heapfy() {
	fmt.Println(ha.Data, "| heapfy ", "| Cnt ", ha.Cnt)

	for i := 1; i <= ha.Cnt/2; {

		maxIndex := i
		if ha.Data[i] < ha.Data[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= ha.Cnt && ha.Data[maxIndex] < ha.Data[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(ha.Data, i, maxIndex)
		i = maxIndex
	}
}

func swap(arr []int, a int, b int) {
	t := arr[a]
	arr[a] = arr[b]
	arr[b] = t
}
