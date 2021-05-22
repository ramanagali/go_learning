package main

import "fmt"

func main() {

	//pointers
	var a int = 12
	var p *int = &a

	fmt.Println(a)
	fmt.Println(p, *p)

	//nil pointer
	var ptr *int
	fmt.Println(ptr)
	println("-----------------")

	//
	type Foo struct {
		bar int
	}

	var foo *Foo
	fmt.Println("Before Assigning memory", foo)

	foo = new(Foo)
	fmt.Println("After Assigning memory", *foo)

	foo.bar = 45
	fmt.Println(foo.bar)
}
