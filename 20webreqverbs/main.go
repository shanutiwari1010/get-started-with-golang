package main

import (
	
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main(){
	fmt.Println("Welcome to web verb video")
	//PerformGetPequest()
	//PerformPostJsonRequest()
	PerformFormPostFormrequest()
}

func PerformGetPequest(){
	const myurl = "http://localhost:8000/get"

	response,err := http.Get(myurl)

	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length : ", response.ContentLength)

	var responsetring strings.Builder
	content,_ :=ioutil.ReadAll(response.Body)
	byteCount, _ := responsetring.Write(content)
	fmt.Println(byteCount)

	fmt.Println(responsetring.String())
	//fmt.Println(content)

	//fmt.Println(string(content))
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
		"cousrsename" : "Lets go with golang",
		"price" : 0,
		"platform" : "learncodeonline.in"

	}
	`)

	response, err:= http.Post(myurl,"application/json",requestBody)
	if err!=nil {
		panic(err)
	}

	defer response.Body.Close()
	content ,_ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformFormPostFormrequest(){
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "Shanu")
	data.Add("lastname" , "Tiwari")
	data.Add("enail", "shanu@qqq")

	response,err := http.PostForm(myurl, data)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()

	content,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
    
}