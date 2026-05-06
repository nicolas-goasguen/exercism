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
}

var ErrListEmpty = errors.New("list is empty")
var ErrListNodeDeletion = errors.New("unable to delete node from list")

func NewNode(v any, prev *Node, next *Node) *Node {
	return &Node{v, prev, next}
}

func NewList(elements ...any) *List {
	if len(elements) == 0 {
		return &List{first: nil}
	}
	first := NewNode(elements[0], nil, nil)
	prev := first
	for _, e := range elements[1:] {
		node := NewNode(e, prev, nil)
		prev.next = node
		prev = prev.next
	}
	return &List{first: first}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) deleteNode(node *Node) bool {
	currentPrev := node.Prev()
	currentNext := node.Next()
	if currentPrev != nil {
		currentPrev.next = currentNext
	} else {
		l.first = currentNext
	}
	if currentNext != nil {
		currentNext.prev = currentPrev
	}
	return true
}

func (l *List) Unshift(v any) {
	first := l.First()
	new := NewNode(v, nil, first)
	if first != nil {
		first.prev = new
	}
	l.first = new
}

func (l *List) Push(v any) {
	last := l.Last()
	new := NewNode(v, last, nil)
	if last != nil {
		last.next = new
	} else {
		l.first = new
	}
}

func (l *List) Shift() (any, error) {
	first := l.First()
	if first == nil {
		return nil, ErrListEmpty
	}
	ok := l.deleteNode(first)
	if !ok {
		return nil, ErrListNodeDeletion
	}
	return first.Value, nil
}

func (l *List) Pop() (any, error) {
	last := l.Last()
	if last == nil {
		return nil, ErrListEmpty
	}
	ok := l.deleteNode(last)
	if !ok {
		return nil, ErrListNodeDeletion
	}
	return last.Value, nil
}

func (l *List) Reverse() {
	last := l.Last()
	current := last
	for current != nil {
		prev := current.prev
		current.prev = current.next
		current.next = prev
		current = current.next
	}
	l.first = last
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	last := l.First()
	if last == nil {
		return nil
	}
	for last.Next() != nil {
		last = last.Next()
	}
	return last
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
			return l.deleteNode(current)
		}
		current = current.Next()
	}
	return false
}
