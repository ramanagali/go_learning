package main

import (
	"fmt"
	"time"
)

func main() {

	//infitinite loop
	for false {
		fmt.Printf("This loop will run forever.\n")
		time.Sleep(time.Second)
	}
	//1
	println("-----------------")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//2
	println("-----------------")
	for i, j := 6, 6; i < 10; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	//3
	println("-----------------")
	i := 50
	for i < 55 {
		fmt.Println(i)
		i++
	}

	//4 netsted loops
	println("-----------------")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i * j)
		}
	}

	// with array loop
	println("-----------------")
	a := []int{1, 2, 3, 4, 5, 6}
	for k, v := range a {
		fmt.Println(k, v)
	}

	//with maps
	cart := make(map[string]int)
	cart = map[string]int{
		"Keyboard": 99,
		"mouse":    50,
		"laptop":   999,
	}
	cart["monitor"] = 499
	for k, v := range cart {
		fmt.Println(k, v)
	}
}
