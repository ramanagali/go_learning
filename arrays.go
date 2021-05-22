package main

import "fmt"

func main() {

	//array
	var amounts [3]int = [3]int{10, 20, 30}
	var amounts2 []int = []int{10, 20, 30, 40}
	amt := []int{1, 2, 3, 4, 5}
	amt[0] = 11
	amt[4] = 55
	fmt.Printf("amounts: %v\n", amounts)
	fmt.Printf("amounts2: %v\n", amounts2)
	fmt.Printf("amt: %v\n", amt)
	fmt.Printf("amt len: %v\n", len(amt))

	//clone
	amt2 := amt
	amt2[0] = 111
	fmt.Printf("amt2: %v\n", amt2)
	fmt.Printf("amt: %v\n", amt)

	//array copy slicking
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[2:]
	d := a[:5]
	e := a[2:7]
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)

	fmt.Printf("slicing : %v\n", a[2:len(a)-1])
	println("-----------------------")

	//multi dimention array
	var idmatrix [][]int = [][]int{
		{11, 12, 13},
		{14, 15, 16},
		{17, 18, 19},
	}
	fmt.Printf("md array: %v\n", idmatrix)
	fmt.Printf("md arr, 0,0th element: %v\n", idmatrix[0][0])

	idmatrix[2][0] = 77
	idmatrix[2][1] = 78
	idmatrix[2][2] = 79
	fmt.Printf("md array: %v\n", idmatrix)
}
