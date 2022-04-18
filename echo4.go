package main

import (
	"flag"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "-", "separator")

func main() {
	//flag.Parse()
	//fmt.Print(strings.Join(flag.Args(),*sep))
	//if !*n {
	//	fmt.Println()
	//}
	//fmt.Println(*n)
	//
	//fmt.Println(gcd(27,81))
	//
	//fmt.Println(popcount.PopCount(12))


	//var l uint8 = 0xff
	//fmt.Println(l)
	//var r uint8 = 0x1
	//fmt.Println(r)
	//var result = l &^ r
	//fmt.Println(result)

	//var l uint8 = 0xff
	//fmt.Println(l<<8)

	//fmt.Println(l>>8)

	//var l int8 = 0x7f
	//fmt.Println(l<<8)

	//fmt.Println(l<<1>>1>>2>>1)
	//fmt.Println(l>>16)
}


func delta(old, new int) int {
	return new - old
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}