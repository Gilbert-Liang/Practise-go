package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	values := []int{2, 3, 9, 7, 1, 8, 12, 24, 13}
	fmt.Println(Sort(values))
}
