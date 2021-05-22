package main

import "fmt"

//structs are value types
//collection of datatypes
type Student struct {
	name     string
	rollNo   int
	subjects []string
}

func main() {
	s1 := Student{
		name:   "Venkat",
		rollNo: 1,
		subjects: []string{
			"Maths",
			"Physics",
			"Chemistry",
		},
	}

	s2 := Student{
		"Manasvi",
		2,
		[]string{
			"English",
			"Art",
			"Computers",
		},
	}

	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.rollNo)
	fmt.Println(s1.subjects[0])

	s1.subjects = append(s1.subjects, "Computers")
	fmt.Println(s1.subjects)
	fmt.Println(s2)

	s3 := s2
	s3.name = "Shanvi"
	s3.rollNo = 3
	fmt.Println(s3)
}
