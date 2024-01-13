package main

import "fmt"

func main()  {
	fmt.Println("Welcome to the class on pointers")

	// var ptr *int
	// fmt.Println("Value of Pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value of actual pinter is", ptr)
	fmt.Println("Value of actual pinter is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is ", myNumber)

	

}