package main

import (
	"fmt"
	"log"
	"os/exec"
)

//run in linux
func main() {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
