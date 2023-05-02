package leetcode

type LRUCache struct {
	head     *Node
	tail     *Node
	cache    map[int]*Node
	capacity int
}

type Node struct {
	next *Node
	prev *Node
	val  int
	key  int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		head:     &Node{},
		tail:     &Node{},
		cache:    make(map[int]*Node, capacity),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.remove(n)
	this.setHead(n)
	return n.val
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.cache[key]
	if ok {
		this.remove(n)
	} else {
		if len(this.cache) == this.capacity {
			delete(this.cache, this.tail.prev.key)
			this.remove(this.tail.prev)
		}
		n = &Node{val: value, key: key}
		this.cache[key] = n
	}
	n.val = value
	this.setHead(n)
}

func (this *LRUCache) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (this *LRUCache) setHead(n *Node) {
	n.next = this.head.next
	n.prev = this.head

	this.head.next.prev = n
	this.head.next = n
}
