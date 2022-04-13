package main

import "fmt"

func main() {
	var values = []int{5, 8, 4, 7, 3 ,9, 1, 11, 6, 0}
	fmt.Println(values)
	quicksort(values,0,len(values)-1)
	fmt.Println(values)
}

func quicksort(values []int, left, right int) {
	l := left
	r := right

	if left >= right {
		return
	}
	pivot := values[(l + r)/2]

	for l < r {
		//fmt.Println("==========")
		for values[l] < pivot {
			l++
		}
		for values[r] > pivot {
			r--
		}

		if l < r {
			temp := values[l]
			values[l] = values[r]
			values[r] = temp
			fmt.Println(values)
		}
		//fmt.Println("==========")
	}

	quicksort(values,left,r)
	quicksort(values,r+1,right)

}
