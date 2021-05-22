package main

import "fmt"

func main() {
	// maps are like dictionaries
	cart := make(map[string]int)
	cart = map[string]int{
		"Keyboard": 99,
		"mouse":    50,
		"laptop":   999,
	}
	//copy
	cart2 := cart

	fmt.Println(cart)
	fmt.Println(len(cart))

	//add
	cart["monitor"] = 499

	//update
	cart["mouse"] = 55
	fmt.Println(cart["mouse"])

	mobile, ok := cart["mobile"]
	_, ok2 := cart["apple"]
	fmt.Println(cart["mobile"])
	fmt.Println(mobile, ok)
	fmt.Println(ok2)

	//delete
	delete(cart, "laptop")

	fmt.Println(cart)
	fmt.Println(cart2)
}
