// Fetch выводит ответ на запрос по заданному URL.
//Измените программу fetch так, чтобы к каждому аргументу URL
// автоматически добавлялся префикс http://
// в случае отсутствия в нем такового.
// Можете воспользоваться функцией strings.HasPrefix.
// go run net/fetch_hw1.8.go http://gopl.io
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	httpScheme = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpScheme) {
			url = httpScheme + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
