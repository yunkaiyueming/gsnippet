package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//GO版本的多态，接口是该方法更具一般性
//此方法的应用在 type A,B都有一个ok()方法，于是定义个m接口有ok()方法，然后写个callOK()里传入m接口类型的方法，就可以根据类型自己调用了
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	gets := []geometry{r, c}
	for _, v := range gets {
		measure(v)
	}
}
