package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int {
		"alice": 31,
		"charlie": 34,
	}

	fmt.Println(ages["alice"])
	fmt.Println(ages["tests"])
	ages["tests"] = ages["alice"] + 2
	delete(ages, "ok")
	fmt.Println(ages["tests"])

	//_ = &ages["alice"] //compile error

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n",name, ages[name])
	}

	var salary map[string]int
	fmt.Println(salary == nil)
	fmt.Println(len(salary) == 0)
	salary = make(map[string]int)
	salary["carol"] = 21

}
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}