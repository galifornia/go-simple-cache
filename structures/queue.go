package structures

import "fmt"

const MAX_QUEUE_SIZE = 10

type Node struct {
	Data string
	Prev *Node
	Next *Node
}
type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
	Size   int
}

type Hash map[string]*Node

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return Queue{Head: head, Tail: tail, Size: MAX_QUEUE_SIZE, Length: 0}
}

func (q *Queue) Display() {
	p := q.Head.Next

	for p != q.Tail {
		fmt.Print(p.Data)
		if p.Next != q.Tail {
			fmt.Print(" --> ")
		}
		p = p.Next
	}
}
