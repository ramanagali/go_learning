package main

import "fmt"

type Processor struct {
	processorName string
	cores         int
}

type Memory struct {
	memoryCapacity int
	memorytType    string
}

//embeding (composition) relationship
//has relationship
type Computer struct {
	Processor
	Memory
	price int
	name  string
}

func main() {
	comp1 := Computer{}
	comp1.name = "Dell Xps"
	comp1.price = 50000
	comp1.processorName = "Intel i7"
	comp1.cores = 8
	comp1.memoryCapacity = 8
	comp1.memorytType = "DDR5"

	comp2 := Computer{
		name: "Macbook pro",
		Processor: Processor{
			processorName: "Intel i9",
			cores:         8,
		},
		Memory: Memory{
			memoryCapacity: 32,
			memorytType:    "DDR5",
		},
		price: 70000,
	}

	fmt.Println(comp1)
	fmt.Println(comp2)
}
