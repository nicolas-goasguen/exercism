// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	sharedWith := name
	if len(sharedWith) < 1 {
		sharedWith = "you"
	}
	res := fmt.Sprintf("One for %v, one for me.", sharedWith)
	return res
}
