package queue

type QueueOnArray struct {
	data      []interface{}
	capactity int
	head      int
	tail      int
}

func NewArrayQueue(n int) *QueueOnArray {
	return &QueueOnArray{
		data:      make([]interface{}, n, n),
		capactity: n,
		head:      0,
		tail:      0,
	}
}

func (aq *QueueOnArray) EnQ(item interface{}) bool {
	// this is a ring
	if aq.head == (aq.tail+1)%aq.capactity {
		return false
	}

	aq.data[aq.tail] = item
	aq.tail = (aq.tail + 1) % aq.capactity
	return true
}

func (aq *QueueOnArray) DeQ() *interface{} {
	if aq.head == aq.tail {
		return nil
	}
	v := aq.data[aq.head]
	aq.head = (aq.head + 1) % aq.capactity
	return &v
}

func (aq *QueueOnArray) IsEmpty() bool {
	return aq.head == aq.tail
}

func (aq *QueueOnArray) Size() int {
	return (aq.tail + aq.capactity - aq.head) % aq.capactity
}
