package main

import (
	"fmt"
	"strconv"
)

type MyStringer interface {
	String() string
}

type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + "°C"
}

type Point struct {
	x int
	y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

var x MyStringer

func main() {
	x = Temp(24)
	fmt.Println(x) // 24 °C

	x = &Point{1, 2}
	fmt.Println(x) // (1,2)ng())
}
