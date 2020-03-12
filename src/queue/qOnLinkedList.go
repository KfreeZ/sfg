package queue

type qOnLinkedList struct {
	head *node
	tail *node
}

type node struct {
	data interface{}
	next *node
}

func NewLinkedlistQueue() *qOnLinkedList {
	return &qOnLinkedList{nil, nil}
}

func (lq *qOnLinkedList) EnQ(item interface{}) bool {
	if lq.head == nil {
		lq.head = &node{data: item, next: nil}
		lq.tail = lq.head
		return true
	}

	lq.tail.next = &node{data: item, next: nil}
	lq.tail = lq.tail.next
	return true

}

func (lq *qOnLinkedList) DeQ() interface{} {
	if lq.head == nil {
		return nil
	}

	v := lq.head.data
	lq.head = lq.head.next
	return v
}
