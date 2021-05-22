package main

import (
	"fmt"
	"os"
)

func main() {
	//defer - exec end of method before retruning val
	println("-----------------")

	println(1)
	defer println(2)
	println(3)
	println(4)

	//ex2 panic - after panic next stmts wont excute
	println("-----------------")
	fmt.Println("start")
	panic("problem")
	fmt.Println("end")

	//recover - after panic can recover from it
	println("-----------------")
	_, err := os.Create("/x/TempFile.txt")
	if err != nil {
		panic(err)
	}
}
