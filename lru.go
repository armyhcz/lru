package lru

// 写入时,先从map中查询,如果能查询,如果能查询到值,则将该值的在List中移动到最前面.
//如果查询不到值,则判断当前map是否到达最大值,如果到达最大值则移除List最后面的值,同时删除map中的值,
//如果map容量未达最大值,则写入map,同时将值放在List最前面.

//读取时,从map中查询,如果能查询到值,则直接将List中该值移动到最前面,返回查询结果.

type LRUCache struct {
	size       int
	capacity   int
	cache      map[string]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key        string
	value      interface{}
	prev, next *DLinkedNode
}

func initDLinkedNode(key string, value interface{}) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func (l *LRUCache) Get(key string) interface{} {
	if _, ok := l.cache[key]; !ok {
		return nil
	}
	node := l.cache[key]
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) Put(key string, value interface{}) {
	if _, ok := l.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		l.cache[key] = node
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			removed := l.removeTail()
			delete(l.cache, removed.key)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	}
}

func (l *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[string]*DLinkedNode{},
		head:     initDLinkedNode("", nil),
		tail:     initDLinkedNode("", nil),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}
