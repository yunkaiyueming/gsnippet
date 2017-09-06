package main

import (
	"context"
	"fmt"
	"time"
)

func contextWithTimeOut() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func contextWithTimeOut2() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}

func contextFindValue() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("k")
	ctx1 := context.WithValue(context.Background(), k, "v1")

	k2 := favContextKey("k2")
	ctx2 := context.WithValue(context.Background(), k2, "v2")

	f(ctx1, k)
	f(ctx1, k2)
	f(ctx2, k)
	f(ctx2, k2)
}

func main() {
	contextWithTimeOut()
	contextWithTimeOut2()
	contextFindValue()
}
