// Echo2 выводит аргументы командной строки
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d %s\n", i, arg)
	}
}
