### 用Go实现stack的问题

	1. 数组实现：用slice做数据结构最简单，因为slice可以自动扩展
	2. 链表实现：用linked list也可以
	3. 需要注意的问题：对于多线程要加锁，可以用sync.mutex 或者是 sync.RWMutex

### 用Java实现stack的问题
	1. 数组实现：java没有go的slice这种数据结构，如果要用object[] 来实现，首先这种结构的大小是固定的，超过capacity的时候需要扩展
	2. 链表实现：
