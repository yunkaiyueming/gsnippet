package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "q" { // break the loop if text == "q"
		fmt.Print("Enter your text: ")
		scanner.Scan()
		text = scanner.Text()
		if text != "q" {
			fmt.Println("Your text was: ", text)
		}
	}
}
