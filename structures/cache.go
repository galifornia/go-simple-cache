package structures

type Cache struct {
	Queue Queue
	Hash  Hash
}

func CreateCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

func (c *Cache) Insert(str string) {
	var node *Node
	if val, ok := c.Hash[str]; ok {
		node = c.RemoveNode(val)
	} else {
		node = &Node{Data: str}
	}
	c.AddNode(node)
	c.Hash[str] = node
}

func (c *Cache) RemoveNode(node *Node) *Node {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	c.Queue.Length--
	delete(c.Hash, node.Data)

	return node
}

func (c *Cache) AddNode(node *Node) {
	node.Prev = c.Queue.Head
	node.Next = c.Queue.Head.Next
	c.Queue.Head.Next.Prev = node
	c.Queue.Head.Next = node

	c.Queue.Length++

	// If we reach max size drop last node
	if c.Queue.Length > c.Queue.Size {
		c.RemoveNode(c.Queue.Tail.Prev)
	}

	c.Hash[node.Data] = node
}
