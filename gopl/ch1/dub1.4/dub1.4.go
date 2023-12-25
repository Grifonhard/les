package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if countLines(f, counts) {
				fmt.Println(arg)
			}
			f.Close()
		}
	}
	for line, n:= range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines (f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	var b = false
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			b = true
		}
	}
	return b
}