package main

import "fmt"

func main() {

	//1st example
	if i := 2; i == 2 {
		fmt.Println("simple if condition")
	}

	// 2nd example
	cart := make(map[string]int)
	cart = map[string]int{
		"Keyboard": 99,
		"mouse":    50,
		"laptop":   999,
	}
	cart["monitor"] = 499

	if _, ok := cart["monitor"]; ok {
		fmt.Println("item exists")
	}

	//3 rd example
	var x = 50
	// var y = 20

	if x == 0 {
		fmt.Println("x is value is 0")
	} else if x > 25 {
		fmt.Println("x is above 25")
	} else if x < 25 {
		fmt.Println("x is below 25")
	} else {
		fmt.Println("else case")
	}

	//Switch
	a, b := 1, 2
	switch a + b {
	case 1:
		fmt.Println("Sum is 1")
	case 2:
		fmt.Println("Sum is 2")
	case 3:
		fmt.Println("Sum is 3")
	default:
		fmt.Println("Printing default")
	}

	//type switch with break
	var j interface{} = 5
	switch j.(type) {
	case int:
		fmt.Println("type is int")
		fmt.Println("type is int")
		break
		fmt.Println("type is int")
	case float64:
		fmt.Println("type is float")
	case string:
		fmt.Println("type is string")
	default:
		fmt.Println("type is default")
	}

	//fallthrough example
	switch 1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Printing default")
	}

}
