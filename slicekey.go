package main

import "fmt"

var m = make(map[string]int)

func main() {
	key1 := []string{"az","li","po","se"}
	fmt.Println(k(key1))

}

func k(list []string) string {
	return fmt.Sprintf("%q",list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}