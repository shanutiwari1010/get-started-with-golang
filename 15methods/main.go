package main

import "fmt"

func main()  {
	fmt.Println("structs in golang")

	shanu := User{"Shanu" , "shanutiwari.work@gmail.com",true ,20}

	fmt.Println(shanu)
	fmt.Printf("shanu details are : %+v\n" , shanu)
	fmt.Printf("Name is %v and email is %v\n" ,shanu.Name,shanu.Email)
	shanu.GetStatus()
	shanu.NewMail()
	fmt.Printf("Name is %v and email is %v\n" ,shanu.Name,shanu.Email)
    
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
	
}

func (u User) GetStatus(){
fmt.Println("Is User active:", u.Status)
}

func(u User) NewMail(){
	u.Email = "test@.gg"
	fmt.Println("Email of this user is ", u.Email)
}