package main

import (
	"fmt"
)

type Factory interface {
	Porduce(string)
}

type Product interface {
	getName() string
}

type ProductA struct {
	Product
	tt string
}

func (this ProductA) getName() string {
	return "A"
}

type ProductB struct {
	Product
	dd int
}

func (this ProductB) getName() string {
	return "B"
}

type ProductC struct {
	Product
	ww bool
}

func (this ProductC) getName() string {
	return "C"
}

type Create struct {
}

func (c *Create) Porduce(pty string) Product {
	if pty == "A" {
		return ProductA{tt: "a_tt"}
	} else if pty == "B" {
		return new(ProductB)
	}
	return new(ProductC)
}

func main() {
	c := new(Create)
	a := []string{"A", "B", "C"}
	for _, p := range a {
		v := c.Porduce(p)
		fmt.Println(v.getName())
	}
}
