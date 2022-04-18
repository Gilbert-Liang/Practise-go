package main

import (
	"fmt"
)

type Node struct {
	x, y int
	u float32
}

func appendInt(x []Node, y ...Node) []Node {
	var z []Node
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]Node, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	for i, i2 := range y {
		z[len(x)+i] = i2
	}
	return z
}

func test(f func([]Node, ...Node) []Node) {
	var x, y []Node
	for i := 0; i < 32; i++ {
		node := new(Node)
		node.x = 1
		node.y = 2
		node.u = 2.0
		y = f(x, *node)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var m []Node
	for i := 0; i < 32; i++ {
		node := new(Node)
		node.x = 1
		node.y = 2
		node.u = 2.0
		m = f(m, *node)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(m), m)
	}
}

func main() {
	test(appendInt)
	//test(append) why error
}