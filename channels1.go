package main

import (
	"fmt"
	"sync"
)

//context of go
//sending & receiving data
//channels are value types
//unbuffered channel - only 1
//
var wg = sync.WaitGroup{}

func main() {
	//buffer is 10 here
	ch := make(chan int, 10)
	wg.Add(2)

	go func(ch <-chan int) {
		defer wg.Done()
		fmt.Println("fetching value from channel..")

		//featch individually
		// fmt.Println(<-ch)
		// fmt.Println(<-ch)

		//loop through channel and fetch
		// for i := range ch {
		// 	fmt.Println(i)
		// }

		//check if value there in channel
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
	}(ch)

	go func(ch chan<- int) {
		defer wg.Done()
		defer close(ch)

		fmt.Println("sending value to channel...")

		//send data 3 times
		ch <- 12
		ch <- 24
		ch <- 36

		// time.Sleep(time.Second)
		// fmt.Println("received back: ", <-ch)

	}(ch)

	wg.Wait()
}
