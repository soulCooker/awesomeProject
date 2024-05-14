package link

/**
哈希表+双向链表实现LRU缓存
**/

type LRUCache struct {
	capacity  int
	bucket    []*CacheNode
	queueHead *CacheNode
	queueTail *CacheNode
	size      int
}

type CacheNode struct {
	Key       int
	Value     int
	Next      *CacheNode
	Prev      *CacheNode
	QueueNext *CacheNode
	QueuePrev *CacheNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		bucket:   make([]*CacheNode, capacity, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	pos := key % this.capacity
	cur := this.bucket[pos]
	for cur != nil {
		if cur.Key == key {
			this.moveToQueueTail(cur)
			return cur.Value
		}
		cur = cur.Next
	}

	return -1
}

func (this *LRUCache) moveToQueueTail(cur *CacheNode) {
	if this.queueTail == cur {
		return
	}

	if this.queueHead == cur {
		this.queueHead = cur.QueueNext
	} else {
		cur.QueuePrev.QueueNext = cur.QueueNext
		cur.QueueNext.QueuePrev = cur.QueuePrev
	}
	cur.QueueNext = nil
	this.queueTail.QueueNext = cur
	cur.QueuePrev = this.queueTail
	this.queueTail = cur
}

func (this *LRUCache) addQueueTail(node *CacheNode) {
	if this.queueTail == nil {
		this.queueHead = node
		this.queueTail = node
		return
	}

	this.queueTail.QueueNext = node
	node.QueuePrev = this.queueTail
	this.queueTail = node
}

func (this *LRUCache) removeQueueHead() {
	removeNode := this.queueHead
	if this.queueTail == this.queueHead {
		this.queueHead = nil
		this.queueTail = nil
	} else {
		this.queueHead = this.queueHead.QueueNext
		this.queueHead.QueuePrev = nil
	}

	pos := removeNode.Key % this.capacity
	if this.bucket[pos] == removeNode {
		this.bucket[pos] = this.bucket[pos].Next
	} else {
		if removeNode.Next != nil {
			removeNode.Next.Prev = removeNode.Prev
		}
		removeNode.Prev.Next = removeNode.Next
	}
}

func (this *LRUCache) Put(key int, value int) {
	pos := key % this.capacity
	cur := this.bucket[pos]
	for cur != nil {
		if cur.Key == key {
			cur.Value = value
			this.moveToQueueTail(cur)
			return
		}
		cur = cur.Next
	}

	cur = &CacheNode{
		Value: value,
		Key:   key,
	}
	if this.bucket[pos] == nil {
		this.bucket[pos] = cur
	} else {
		cur.Next = this.bucket[pos]
		this.bucket[pos].Prev = cur
		this.bucket[pos] = cur
	}
	this.addQueueTail(cur)
	if this.size < this.capacity {
		this.size++
	} else {
		this.removeQueueHead()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
