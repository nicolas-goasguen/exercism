package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value any
	prev  *Node
	next  *Node
}

type List struct {
	first *Node
	last  *Node
}

var ErrListEmpty = errors.New("list is empty")

func NewNode(v any) *Node {
	return &Node{v, nil, nil}
}

func NewList(elements ...any) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func (n *Node) Prev() *Node {
	return n.prev
}

func linkNodes(prev *Node, next *Node) {
	if next != nil {
		next.prev = prev
	}
	if prev != nil {
		prev.next = next
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}

func (l *List) Unshift(v any) {
	newNode := NewNode(v)
	first := l.First()

	l.first = newNode
	if first == nil {
		l.last = newNode
	} else {
		linkNodes(newNode, first)
	}
}

func (l *List) Push(v any) {
	newNode := NewNode(v)
	last := l.Last()

	if last == nil {
		l.first = newNode
	} else {
		linkNodes(last, newNode)
	}
	l.last = newNode
}

func (l *List) Shift() (any, error) {
	first := l.First()
	if first == nil {
		return nil, ErrListEmpty
	}
	next := first.Next()
	if next != nil {
		next.prev = nil
	} else {
		l.last = nil
	}
	l.first = next

	return first.Value, nil
}

func (l *List) Pop() (any, error) {
	last := l.Last()
	if last == nil {
		return nil, ErrListEmpty
	}
	prev := last.Prev()
	if prev != nil {
		prev.next = nil
	} else {
		l.first = nil
	}
	l.last = prev

	return last.Value, nil
}

func (l *List) Reverse() {
	first := l.First()
	if first == nil {
		return
	}
	last := l.Last()

	current := first
	for current != nil {
		prev := current.Prev()
		next := current.Next()
		current.next = prev
		current.prev = next
		current = next
	}

	l.first = last
	l.last = first
}

func (l *List) Count() int {
	current := l.First()
	count := 0
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (l *List) Delete(v any) bool {
	current := l.First()
	for current != nil {
		if current.Value == v {
			prev := current.Prev()
			next := current.Next()
			if prev == nil {
				l.first = next
			}
			if next == nil {
				l.last = prev
			}
			linkNodes(prev, next)
			return true
		}
		current = current.Next()
	}
	return false
}
