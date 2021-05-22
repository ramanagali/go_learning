package main

import (
	"fmt"
	"sync"
	"time"
)

//sepcial type of functions to execute indipendently
//light weight threads

func main() {
	//Wait group
	var wg = sync.WaitGroup{}

	//1st go routine
	wg.Add(1)
	go writeMsg(&wg)

	//2nd go routine
	wg.Add(1)
	msg := "Hello User"
	go func(str string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Anonymous func", str)
	}(msg, &wg)
	//to test Data race
	msg = "Hello GVR"

	//3rd example
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	//program will wait until all go routines done
	wg.Wait()
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker starting ", id)
	time.Sleep(time.Second)
	fmt.Println("worker ended ", id)
}

func writeMsg(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("writeMsg func")
}
