package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	_ "os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

var client *rpc.Client
var args Args

func main() {
	//fmt.Println(os.Args)

	//	if len(os.Args) != 2 {
	//		fmt.Println("Usage: ", os.Args[0], "127.0.0.1:1234")
	//		log.Fatal(1)
	//	}
	//service := os.Args[1]

	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args = Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	syncCall()
	//AsyncCall()
}

func syncCall() {
	//sync call
	var quot Quotient
	var err error
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}

func AsyncCall() {
	// Asynchronous call
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	fmt.Println(replyCall)
	// check errors, print, etc.
}
