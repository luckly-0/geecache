package lru

import "container/list"

type Cache struct {
	maxBytes  int64                         // 允许的最大内存
	nbytes    int64                         // 已使用的内存
	ll        *list.List                    //双向链表
	cache     map[string]*list.Element      // map
	OnEvicted func(key string, value Value) //是某条记录被移除时的回调函数
}

type entry struct { // 代表双向链表的数据类型
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache { // 创建
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) { // 查找
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)    // 移动到队尾
		kv := ele.Value.(*entry) // 找到值
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() { // 缓存淘汰
	ele := c.ll.Back() // 取到队首节点，从链表中删除
	if ele != nil {    // 非空
		c.ll.Remove(ele) // 移除最近最少访问
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value) // 回调
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest() //如果超过了设定的最大值 c.maxBytes，则移除最少访问的节点。
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
