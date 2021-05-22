package main

import (
	"fmt"
	"time"
)

//select similar to switch
//used select from multiple channels

func main() {
	r1 := make(chan string)
	r2 := make(chan string)

	go portal1(r1)
	go portal2(r2)

	{
		select {
		case op1 := <-r1:
			fmt.Println(op1)
		case op2 := <-r2:
			fmt.Println(op2)
		}
	}
}

func portal1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Welecom to Channel 1"
}

func portal2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Welecom to Channel 2"
}
