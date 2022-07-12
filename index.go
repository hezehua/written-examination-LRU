package LRU

//实现一个缓存Cache， 缓存Cache会在其已满时驱逐最近最少使用的数据
//例如：
//缓存大小 ： Size 3
//输入: Data ("key1", "1"), Data ("key2", "2"), Data ("key3", "3"),  Get "key1",  Put ("key4", "4")
//输出: [("key4", "4"), ("key1", "1"), ("key3", "3")]

type LRU struct {
	cap int
	data map[string]dataNode
	cacheLinkHead *cacheNode
	cacheLinkTail *cacheNode
}

type dataNode struct {
	val string
	point *cacheNode
}

type cacheNode struct {
	key string
	pre *cacheNode
	next *cacheNode
}

//设置缓存
func (l *LRU) set(key, val string) bool {
	if l.cap == 0 {
		return false
	}
	dataLen := len(l.data)
	if dataLen == 0 {
		l.data = make(map[string]dataNode)
	}
	ret , ok := l.data[key]
	if ok {
		//有该缓存，直接更新
		ret.val = val
		l.cutNode(ret.point)
		l.inertToHead(ret.point)
	} else if dataLen < l.cap {
		//无该缓存，且还有容量
		nodeTmp := cacheNode{
			key:  key,
		}
		l.inertToHead(&nodeTmp)
		l.data[key] = dataNode{
			val:val,
			point: &nodeTmp,
		}
	} else {
		//无该缓存，且没有容量
		nodeTmp := cacheNode{
			key:  key,
		}
		oldNode, ok := l.removeTail()
		if ok {
			delete(l.data, oldNode.key)
		}
		l.data[key] = dataNode{
			val:val,
			point: &nodeTmp,
		}
		l.inertToHead(&nodeTmp)
	}
	return true
}

//获取缓存
func (l *LRU) get(key string) (string, bool) {
	ret , ok := l.data[key]
	if !ok {
		return "", false
	}
	l.cutNode(ret.point)
	l.inertToHead(ret.point)
	return ret.val, true
}

//插入到头部
func (l *LRU) inertToHead(node *cacheNode)  {
	node.pre = nil
	if l.cacheLinkHead != nil {
		node.next = l.cacheLinkHead
		if node.next != nil {
			node.next.pre = node
		}
	}
	l.cacheLinkHead = node
	if l.cacheLinkTail == nil {
		l.cacheLinkTail = node
	}
}

//将节点切出
func (l *LRU) cutNode(p *cacheNode) {
	if p.next != nil {
		p.next.pre = p.pre
	}
	if p == l.cacheLinkHead {
		l.cacheLinkHead = nil
	}
	if p.pre != nil {
		p.pre.next = p.next
		if p == l.cacheLinkTail {
			l.cacheLinkTail = p.pre
		}
	}
}

//移除尾节点
func (l *LRU) removeTail() (*cacheNode, bool) {
	tail := l.cacheLinkTail
	if tail == nil {
		return nil, false
	}
	if l.cacheLinkTail.pre != nil {
		l.cacheLinkTail.pre.next = nil
		l.cacheLinkTail = l.cacheLinkTail.pre
	} else {
		l.cacheLinkTail = nil
	}
	return tail, true
}