package geecache

import (
	"container/list"
)

type Cache struct {
	maxBytes  int64
	nBytes    int64
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		// 如果键对应的链表节点存在，则将对应节点移动到队尾，并返回查找到的值。
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	// c.ll.Back() 取到队首节点，从链表中删除。
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		// delete(c.cache, kv.key)，从字典中 c.cache 删除该节点的映射关系。
		delete(c.cache, kv.key)
		// 更新当前所用的内存 c.nbytes
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			// 调用回调函数
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		// 如果键存在，则更新对应节点的值，并将该节点移到队尾
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		// 不存在则是新增场景，首先队尾添加新节点 &entry{key, value},
		// 并字典中添加 key 和节点的映射关系。
		ele := c.ll.PushFront(&entry{key: key, value: value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	// 更新 c.nbytes，如果超过了设定的最大值 c.maxBytes，则移除最少访问的节点。
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
