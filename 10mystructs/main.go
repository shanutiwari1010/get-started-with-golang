package main

import "fmt"

func main()  {
	fmt.Println("structs in golang")

	shanu := User{"Shanu" , "shanutiwari.work@gmail.com",true ,20}

	fmt.Println(shanu)
	fmt.Printf("shanu details are : %+v\n" , shanu)
	fmt.Printf("Name is %v and email is %v" ,shanu.Name,shanu.Email)

}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

