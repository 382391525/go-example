package lcache

import "container/list"

type fifo struct {
	ll        *list.List
	cache     map[string]*list.Element
	maxBytes  int
	useBytes  int
	onEvicted func(key string, value interface{})
}

func NewFifo(maxBytes int, onEvicted func(key string, value interface{})) *fifo {
	return &fifo{
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		maxBytes:  maxBytes,
		onEvicted: onEvicted,
	}
}

func (f *fifo) Set(key string, value interface{}) {
	if e, ok := f.cache[key]; ok {
		f.ll.MoveToBack(e)
		en := e.Value.(*entry)
		f.useBytes = f.useBytes - CalcLen(en.value) + CalcLen(value)
		en.value = value
		return
	}
	en := &entry{
		key:   key,
		value: value,
	}
	e := f.ll.PushBack(en)
	f.cache[key] = e
	f.useBytes += en.Len()
	if f.maxBytes > 0 && f.useBytes > f.maxBytes {
		f.DelOldest()
	}
}

func (f *fifo) Get(key string) interface{} {
	if element, ok := f.cache[key]; ok {
		return element.Value.(*entry).value
	}
	return nil
}

func (f *fifo) Del(key string) {
	if element, ok := f.cache[key]; ok {
		f.removeElement(element)
	}
}

func (f *fifo) removeElement(e *list.Element) {
	if e == nil {
		return
	}
	f.ll.Remove(e)
	en := e.Value.(*entry)
	f.useBytes -= en.Len()
	delete(f.cache, en.key)
	if f.onEvicted != nil {
		f.onEvicted(en.key, en.value)
	}
}

func (f *fifo) DelOldest() {
	f.removeElement(f.ll.Front())
}

func (f *fifo) Len() int {
	return f.ll.Len()
}
