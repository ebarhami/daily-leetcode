type Node struct {
    key int
    val int
    next *Node
    prev *Node
}

type LRUCache struct {
    head *Node // put least recently used in head
    tail *Node
    nodeByKey map[int]*Node
    size int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        size: capacity,
        nodeByKey: make(map[int]*Node),
    }
}

func (this *LRUCache) insert(key, value int) *Node {
    node := &Node{
        key: key,
        val: value,
        prev: this.tail,
    }
    if this.head == nil {
        this.head = node
    }

    if this.tail != nil {
        this.tail.next = node
    }
    this.tail = node

    return node
}

func (this *LRUCache) delete(node *Node) {
    if node.next != nil {
        node.next.prev = node.prev
    } else { // node at tail
        this.tail = node.prev
    }

    if node.prev != nil {
        node.prev.next = node.next
    } else { // node at head
        this.head = node.next
    }
}

func (this *LRUCache) Get(key int) int {
    if node, exist := this.nodeByKey[key]; exist {
        this.delete(node)
        nodeNew := this.insert(key, node.val)
        this.nodeByKey[key] = nodeNew
        return node.val
    }

    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    node := &Node{val: value, key:key}
    if _, exist := this.nodeByKey[key]; exist {
        node = this.nodeByKey[key]
        this.delete(node)
    }else if this.size == len(this.nodeByKey) {
        delete(this.nodeByKey, this.head.key)
        this.delete(this.head)
    }
    nodeNew := this.insert(key, value)
    this.nodeByKey[key] = nodeNew
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */