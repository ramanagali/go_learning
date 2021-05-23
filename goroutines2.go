package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i <= value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Concurrenlty with Go routines")
	go compute(5)
	go compute(5)

	fmt.Scanln()
}
