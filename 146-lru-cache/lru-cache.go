type DoubleLinkedList struct {
    next *DoubleLinkedList
    prev *DoubleLinkedList
    val int
    key int
}

type LRUCache struct {
    head *DoubleLinkedList
    tail *DoubleLinkedList
    capacity int
    lookup map[int]*DoubleLinkedList
}


func Constructor(capacity int) LRUCache {
    h, t := &DoubleLinkedList{}, &DoubleLinkedList{}
    h.next = t
    t.prev = h
    return LRUCache{
        capacity: capacity,
        head: h,
        tail: t,
        lookup: make(map[int]*DoubleLinkedList),
    }
}

/*
head -> b
head -> a -> b
*/
func (this *LRUCache) insert(key, value int) {
    if len(this.lookup) >= this.capacity {
        this.remove(this.tail.prev.key)
    }
    node := &DoubleLinkedList{
        val: value,
        key: key,
    }
    node.next = this.head.next
    node.prev = this.head
    this.head.next = node
    node.next.prev = node
    this.lookup[key] = node
}

/*
    a -> b -> c
*/
func (this *LRUCache) remove(key int) {
    node := this.lookup[key]
    if node == nil {
        return
    }
    node.prev.next = node.next
    node.next.prev = node.prev
    delete(this.lookup, key)
}


func (this *LRUCache) Get(key int) int {
    node := this.lookup[key]
    if node == nil {
        return -1
    }
    this.remove(key)
    this.insert(key, node.val)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    node := this.lookup[key]
    if node != nil {
        this.remove(key)
    }
    this.insert(key, value)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */