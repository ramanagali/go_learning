package main

import "fmt"

func main() {

	//boolean
	foo := true
	fmt.Printf("%v, %T \n", foo, foo)

	i := 20
	j := 3

	//arthematic operators
	fmt.Printf("%v, %T \n", i, j)
	println(i + j)
	println(i - j)
	println(i * j)
	println(i / j)
	println(i % j)
	println("------")

	//binary operators
	println(i & j)
	println(i | j)
	println(i ^ j)
	println(i &^ j)
	println("------")

	//float
	a := 3.14
	b := 5.6

	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	println("------")

	//complex types
	var c complex128 = complex(2, 4) //1 + 2i
	d := 1 + 6i
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", real(c), real(c))
	fmt.Printf("%v, %T \n", imag(c), imag(c))

	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)

	s := "This is String"
	s1 := 'A'
	fmt.Printf("%v, %T \n", s, s)
	fmt.Printf("%v, %T \n", string(s[1]), s)
	fmt.Printf("%v, %T \n", s1, s1)

}
