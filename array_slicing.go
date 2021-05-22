package main

import "fmt"

func main() {

	//array
	var a []int = []int{1, 2, 3}
	var b []int = a
	var c []int = make([]int, 3, 10)

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 0
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(len(c))
	fmt.Println(cap(c))

	//inject
	fmt.Println(append(a, 5))
	fmt.Println(append(a[1:], 5))

	d := append(b, a...)
	fmt.Println(d)
	// fmt.Printf("amt len: %v\n", len(a))
}
