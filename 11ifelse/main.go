package main

import "fmt"

func main()  {
	fmt.Println("If else in golang")

	logincount := 10
	var result string
    
	if logincount<10{
		result =  "Regular user"
	}else if logincount>10 {
		result = "Watch out"
	}else{
		result = "exatly 10 login count"
	}
	fmt.Println(result)

	if num := 3 ; num <10 {
		fmt.Println("less thn 10")
	}else{
		fmt.Println("not less than 10")
	}
	
}