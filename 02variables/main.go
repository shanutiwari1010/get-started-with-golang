package main

import "fmt"

func main() {
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type:   %T \n", isLoggedIn)

	var smallVal int = 256 
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type:   %T \n", smallVal)

	var smallFloat float64= 255.345354645677
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type:   %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	var str string
	fmt.Println(str)
	fmt.Printf("Variable is of type : %T \n", str)

	numberOfUser := 300000
    fmt.Println(numberOfUser)
}

