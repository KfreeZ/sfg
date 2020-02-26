package stack

import "sync"

// ArrayStack is a stack based on array
// which means the stack is a continious memory space
// we need tp dynamically extend the memory to avoid stack overload
type ArrayStack struct {
	data []interface{}
	top  int
	//considering the concurrent, introduce the RWMutex
	rwLock sync.RWMutex
}

//NewArrayStack constructor
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

//IsEmpty return true is the stack is empty
func (as *ArrayStack) IsEmpty() bool {
	return as.top < 0
}

//Push stack push
func (as *ArrayStack) Push(item interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top++
	}

	if as.top > len(as.data)-1 {
		as.data = append(as.data, item)
	} else {
		as.data[as.top] = item
	}
}

//Pop stack pop
func (as *ArrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}

	v := as.data[as.top]
	as.top--
	return v

}

//Top stack top
func (as *ArrayStack) Top() interface{} {
	if as.IsEmpty() {
		return nil
	}
	return as.data[as.top]
}

//Flush stack flush
func (as *ArrayStack) Flush() {
	as.top = -1
}
