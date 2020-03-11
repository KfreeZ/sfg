package stack

type node struct {
	data interface{}
	next *node
}

type ArrayOnLinkedlist struct {
	top *node
}

func NewLinkedlistArray() *ArrayOnLinkedlist {
	return &ArrayOnLinkedlist{nil}
}

func (lla *ArrayOnLinkedlist) IsEmpty() bool {
	return lla.top == nil
}

func (lla *ArrayOnLinkedlist) Top() interface{} {
	if lla.IsEmpty() {
		return nil
	}
	return lla.top.data
}

func (lla *ArrayOnLinkedlist) Push(item interface{}) {
	lla.top = &node{data: item, next: lla.top}
}

func (lla *ArrayOnLinkedlist) Pop() interface{} {
	if lla.IsEmpty() {
		return nil
	}
	data := lla.Top()
	lla.top = lla.top.next
	return data
}

func (lla *ArrayOnLinkedlist) Flush() {
	lla.top = nil
}
