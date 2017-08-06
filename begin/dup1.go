// Dup1 выводит текст каждой строки, которая появляется в
// стандартном вводе более одного раза, а также количество
// ее появлений.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// считывает очередную строку и удаляет завершающий символ новой строки
	for input.Scan() { // ctrl+Z->Enter чтобы завершить ввод
		counts[input.Text()]++
	}
	// Примечание: игнорируем потенциальные
	// ошибки из input.Err()

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
