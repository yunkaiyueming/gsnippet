package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		rawline, _, _ := r.ReadLine()
		line := string(rawline)
		if line == "q" || line == "e" {
			break
		}

		fmt.Println(line)
	}
}
