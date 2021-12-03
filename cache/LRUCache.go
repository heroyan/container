package cache

type LRUCache struct {
	capacity int
	cache *LinkListMap
}

type LinkItem struct {
	key int
	value int
	prev *LinkItem
	next *LinkItem
}

type LinkList struct {
	head *LinkItem
	tail *LinkItem
	len int
}

func (this *LinkList) insert(item *LinkItem)  {
	item.next = this.head.next
	item.prev = this.head
	this.head.next.prev = item
	this.head.next = item
}

func (this *LinkList) delete(item *LinkItem) {
	item.prev.next = item.next
	item.next.prev = item.prev
}

func (this *LinkList) deleteTail() *LinkItem {
	prev := this.tail.prev
	this.delete(this.tail.prev)
	return prev
}

type LinkListMap struct {
	list *LinkList
	cache map[int]*LinkItem
	listLen int
}

func NewList() *LinkList  {
	head := &LinkItem{
		prev: nil,
	}
	tail := &LinkItem{
		next: nil,
	}
	head.next = tail
	tail.prev = head
	return &LinkList{
		head: head,
		tail: tail,
		len: 0,
	}
}

func NewListMap() *LinkListMap  {
	return &LinkListMap{
		list: NewList(),
		cache: map[int]*LinkItem{},
		listLen: 0,
	}
}

func (this *LinkListMap) get(key int) *LinkItem  {
	return this.cache[key]
}

func (this *LinkListMap) insert(item *LinkItem)  {
	// 头查
	this.list.insert(item)
	this.listLen++
	this.cache[item.key] = item
}

func (this *LinkListMap) delete(item *LinkItem)  {
	this.list.delete(item)
	this.listLen--
	delete(this.cache, item.key)
}

func (this *LinkListMap) deleteTail()  {
	item := this.list.deleteTail()
	this.listLen--
	delete(this.cache, item.key)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache: NewListMap(),
	}
}

func (this *LRUCache) Get(key int) int {
	item := this.cache.get(key)
	if item == nil {
		return -1
	}

	// 删除之后插入到头部
	this.cache.delete(item)
	this.cache.insert(item)

	return item.value
}


func (this *LRUCache) Put(key int, value int)  {
	item := this.cache.get(key)
	if item == nil {
		// 新的item
		if this.cache.listLen >= this.capacity {
			// 删除尾部的
			this.cache.deleteTail()
		}
		item = &LinkItem{
			key: key,
			value: value,
		}
		this.cache.insert(item)
		return
	}
	this.cache.delete(item)
	item.value = value
	this.cache.insert(item)

	return
}

func (this *LRUCache) Result() map[int]*LinkItem {
	return this.cache.cache
}
