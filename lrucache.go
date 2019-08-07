package LruCache

//双向链表
type Element struct {
	prev  *Element
	next  *Element
	key   interface{}
	value interface{}
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

//HashMap
type LruCache struct {
	cache    map[interface{}]*Element
	head     *Element
	tail     *Element
	capacity int
}

func New(capacity int) *LruCache {
	return &LruCache{make(map[interface{}]*Element), nil, nil, capacity}
}

func (lc *LruCache) Put(key interface{}, value interface{}) {
	if e, ok := lc.cache[key]; ok {
		e.value = value
		lc.refresh(e)
		return
	}

	if lc.capacity == 0 {
		return
	} else if len(lc.cache) >= lc.capacity {
		delete(lc.cache, lc.tail.key)
		lc.tail.key = key
		lc.tail.value = value
		lc.cache[key] = lc.tail
		lc.refresh(lc.tail)
		return
	}

	e := &Element{
		prev:  nil,
		next:  lc.head,
		key:   key,
		value: value,
	}
	lc.cache[key] = e
	if len(lc.cache) > 1 {
		lc.head.prev = e
	} else {
		lc.tail = e
	}
	lc.head = e
}

func (lc *LruCache) Get(key interface{}) (interface{}, bool) {
	if e, ok := lc.cache[key]; ok {
		lc.refresh(e)
		return e.value, ok
	}
	return nil, false
}

func (lc *LruCache) Delete(key interface{}) {
	if e, ok := lc.cache[key]; ok {
		delete(lc.cache, key)
		lc.remove(e)
	}
}

func (lc *LruCache) Head() *Element {
	return lc.head
}

func (lc *LruCache) Tail() *Element {
	return lc.tail
}

func (lc *LruCache) Len() int {
	return len(lc.cache)
}

func (lc *LruCache) Capacity() int {
	return lc.capacity
}

//更新缓存链表中元素位置
func (lc *LruCache) refresh(e *Element) {
	if e.prev != nil {
		e.prev.next = e.next
		if e.next == nil {
			lc.tail = e.prev
		} else {
			e.next.prev = e.prev
		}

		e.prev = nil
		e.next = lc.head
		lc.head.prev = e
		lc.head = e
	}
}

func (lc *LruCache) remove(e *Element) {
	if e.prev == nil {
		lc.head = e.next
	} else {
		e.prev.next = e.next
	}

	if e.next == nil {
		lc.tail = e.prev
	} else {
		e.next.prev = e.prev
	}
}
