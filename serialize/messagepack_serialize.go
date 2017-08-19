package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/msgpack"
)

func main() {
	ExampleMarshal()
}

func ExampleMarshal() {
	type Item struct {
		Foo string
	}
	data := &Item{Foo: "bar"}

	b, err := msgpack.Marshal(data)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("msgpack===>", len(b), string(b))
	}

	b1, err := json.Marshal(data)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("json===>", len(b1), string(b1))
	}

	network := new(bytes.Buffer)
	gobenc := gob.NewEncoder(network)
	gobenc.Encode(data)
	fmt.Println("gob===>", network.Len(), string(network.Bytes()))

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item)

	json.Unmarshal(b1, &item)
	fmt.Println(item)

	gobdec := gob.NewDecoder(network)
	gobdec.Decode(&item)
	fmt.Println(item)
}
