package main

import (
	"bufio"
	"fmt"
	"os"
)

type dupInfo struct {
	Text     string
	Filename string
}

func main() {
	counts := make(map[dupInfo]int)
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
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s:%s\n", n, line.Filename, line.Text)
		}
	}
}

func countLines(f *os.File, counts map[dupInfo]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		info := dupInfo{
			Text:     input.Text(),
			Filename: f.Name(),
		}
		counts[info]++
	}
}
