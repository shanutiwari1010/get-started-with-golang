package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to vedio on slices")

	var fruitlist = [] string {"apple", "tomato", "peach"}
	fmt.Printf("Type of fruitlist is %T \n", fruitlist)

	fruitlist = append(fruitlist, "mango", "banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)

	highScores[0] =234
	highScores[1] =945
	highScores[2] =465
	highScores[3] =867
	// highScores[4] =777
    highScores = append(highScores, 444,726,265)

	sort.Ints(highScores)
    
    fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = [] string {"reactjs","javascript" , "swift" ,"python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[ :index], courses[index+1 : ]... )
	fmt.Println(courses)

}