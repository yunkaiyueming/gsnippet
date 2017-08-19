package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		txt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input err", err.Error())
		}

		txt = strings.TrimSpace(txt)
		if txt == "q" {
			fmt.Println("bye!!")
			return
		}
	}
}
