package main

import "fmt"

func main()  {
	fmt.Println("Functions in golang")
	greeter()
   
	result := sum(10, 2)
	fmt.Println("Result is ", result)

	result1,msg:= adder(2,5,6,2)
	fmt.Println(result1)
	fmt.Println(msg)
}

func adder(values... int) (int,string){
	total := 0

	for _,val := range  values{
		total += val
	}
	return total, "Shanuuuu"
}
func greeter(){
	fmt.Println("Namstey from golang")
}
func greeterTwo(){
	fmt.Println("Another Method")
}
func sum(a int ,b int) int{
	return a+b
}