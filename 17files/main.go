package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Welcome to files in golang")
	content := "this needs to go in a file - LearnCodeOnline.in"

	file , err := os.Create("./mylcogofiles.txt")

	checkNilErr(err)

	length,err := io.WriteString(file,content)
	checkNilErr(err)

	fmt.Println("lenght is :", length)

	defer file.Close()
	readFile("./mylcogofiles.txt")

}

func readFile(filename string){
	databyte, err :=ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error){
	if err!=nil {
		panic(err)
    }   
}