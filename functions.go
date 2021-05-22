package main

import "fmt"

type rectangle struct {
	width, height int
}

//method for rectangle type
//rectange as input, method for rectangle type
func (r rectangle) area() int {
	return r.height * r.width
}

func main() {
	//pass the ref
	msg := "Hello GVR"
	writeMessage(&msg)
	fmt.Println(msg)

	//multi parameters
	fmt.Println("sum", sum(1, 2, 3))
	fmt.Println("sum", sum(1, 2, 3, 4))

	//multiple return values
	value, error := divide(3, 2)
	if error != nil {
		fmt.Println((error))
	}
	fmt.Println(value)

	//anonymouns function
	myfun := func() {
		fmt.Println("this is anonymouns function")
	}
	myfun()

	//call method
	area := rectangle{20, 10}.area()
	fmt.Println("Area of Rectangle", area)

}

//multiple parameters
func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

//default
func writeMessage(msg *string) {
	*msg = *msg + " Welcome"
	fmt.Println(*msg)
}

//muilte return
func divide(a, b float64) (float64, error) {
	if b == 0.0 || b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
