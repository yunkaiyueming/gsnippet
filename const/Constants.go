package main

import "fmt"
import "math"

//const identifier [type] = value
const s string = "constant"
const (
	start = iota
	second
	thrid
	forth
	fiveth
)

const a, b, c, d = 1, "2", "3", 4

func main() {
	fmt.Println(start, second, thrid, forth, fiveth)
	fmt.Println(a, b, c, d)

	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	const_demo1()
}

func calc() {
	const t = 1 << 2
}

func const_demo1() {
	const (
		h = iota //h=0
		i = 100  //i=100
		j        //j=100
		k = iota //k=3
	)
	fmt.Println(h, i, j, k)

	const (
		bit00 uint32 = 1 << iota //bit00=1,1移一位
		bit01                    //bit01=2
		bit02                    //bit02=4
	)
	fmt.Println(bit00, bit01, bit02)
}
