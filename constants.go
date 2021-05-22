package main

import "fmt"

const User string = "Venkat"
const Product string = "Nikon"

const (
	User2    string = "Venkat"
	Product2 string = "Nikon"
)

const (
	c1 = iota
	c2
	c3
	_
	c5
)

func main() {
	const i int = 12
	const j float32 = 3.14
	const k string = "Kite"
	const l bool = true
	var a int = 13

	fmt.Println(i + a)
	fmt.Println(User)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c5)
}
