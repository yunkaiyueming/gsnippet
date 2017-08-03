package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}
type Q struct {
	X, Z *int32
	Name string
}

func main() {
	var network bytes.Buffer

	// Encode (send) the value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	} else {
		fmt.Println("encode end data:", network)
	}

	// Decode (receive) the value.
	var q Q
	dec := gob.NewDecoder(&network)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	fmt.Println(q)
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Z)
}
