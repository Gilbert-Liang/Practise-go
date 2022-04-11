package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		sin := make(map[string]int)
		countLines(os.Stdin, sin)
		counts["stdin"] = sin
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			fin := make(map[string]int)
			countLines(f, fin)
			f.Close()
			counts[arg] = fin
		}
	}
	for f, sum := range counts {
		fmt.Printf("-------------\n%s\n===============\n",f)
		for line, n := range sum {
			fmt.Printf("t%d\t%s\n",n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}