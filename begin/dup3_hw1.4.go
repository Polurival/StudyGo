package main

// Измените программу dup3 так, чтобы она выводила имена всех
// файлов, в которых найдены повторяющиеся строки

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		counts := make(map[string]int)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		for _, n := range counts {
			if n > 1 {
				fmt.Printf("%s\n", filename)
				break
			}
		}
	}
}
