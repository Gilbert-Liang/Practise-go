package main

import "fmt"

func main() {
	var x, y int
	var p = f()
	fmt.Println(&x == &x, &x == &y, &x == nil, *p)

	v := 1
	incr(&v)
	fmt.Println(incr(&v))

}


func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++

	return *p
}


