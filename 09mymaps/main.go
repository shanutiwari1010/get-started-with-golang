package main

import (
	"fmt"
	
)

func main()  {
	fmt.Println("maps in golang")

	languages := make(map[string]string)
    languages["JS"] = "javaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Lists of all languages : " , languages)
	fmt.Println("Js shorts for :", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Lists of all languages : " , languages)

    for _, value := range languages{
		fmt.Printf("for key v , value is %v \n",value)
	}

}