package pov

import (
	"slices"
)

type Tree struct {
	value    string
	parent   *Tree
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	newTree := &Tree{value: value, children: children}
	for _, c := range newTree.children {
		c.parent = newTree
	}
	return newTree
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

func (tr *Tree) Parent() *Tree {
	return tr.parent
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions
// FromPov returns the pov from the node specified in the argument.

func (tr *Tree) findNode(value string) *Tree {
	if tr.Value() == value {
		return tr
	}
	for _, c := range tr.Children() {
		n := c.findNode(value)
		if n != nil {
			return n
		}
	}
	return nil
}

func (tr *Tree) rerootFrom(from *Tree) *Tree {
	prevParent := tr.Parent()
	value := tr.value

	var newChildren []*Tree

	for _, c := range tr.Children() {
		if c != from {
			newChildren = append(newChildren, c.rerootFrom(tr))
		}
	}

	if prevParent != nil && prevParent != from {
		newChildren = append(newChildren, prevParent.rerootFrom(tr))
	}

	return New(value, newChildren...)
}

func (tr *Tree) FromPov(from string) *Tree {
	node := tr.findNode(from)
	if node == nil {
		return nil
	}
	return node.rerootFrom(nil)
}

func (tr *Tree) getPathTo(to string, path []string) []string {
	value := tr.Value()
	newPath := append(path, value)

	if value == to {
		return newPath
	}

	for _, c := range tr.Children() {
		nextPath := c.getPathTo(to, newPath)
		if len(nextPath) > 0 {
			return nextPath
		}
	}

	return nil
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	fromPath := tr.getPathTo(from, nil)
	if fromPath == nil {
		return nil
	}
	toPath := tr.getPathTo(to, nil)
	if toPath == nil {
		return nil
	}

	// after getting the paths from root to A to B,
	// we will remove the path until the last common node
	// to deduce the real path from A to B
	prevIdx := 0
	for i := 0; i < len(fromPath); i++ {
		if i >= len(toPath) || fromPath[i] != toPath[i] {
			break
		}
		prevIdx = i
	}

	commonNode := fromPath[prevIdx]

	// reverse the path from the common point to start from A
	fromPath = fromPath[prevIdx+1:]
	slices.Reverse(fromPath)

	// get the path from common point to B
	toPath = toPath[prevIdx+1:]

	// generate the real path from A to B
	fromToPath := append(fromPath, commonNode)
	return append(fromToPath, toPath...)
}
