// Fetch выводит ответ на запрос по заданному URL.
// Измените программу fetch так,
// чтобы она выводила код состояния HTTP, содержащийся в resp. Status.
// go run net/fetch_hw1.9.go http://gopl.io
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("http status code: %s", resp.Status)
	}
}
