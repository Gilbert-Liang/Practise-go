package main

import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove1(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}


func main() {
	s := []int{5, 6, 7, 8, 9}
	s = remove1(s, 2)
	fmt.Println(s) // "[5 6 8 9]"
	s = remove1(s, 2)
	fmt.Println(s) // "[5 6 9 8]"
}