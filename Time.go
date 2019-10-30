package main

import "fmt"
import "time"

func timeTry(){
	fmt.Println("Time is : ", time.Now())
	fmt.Println("Time2 is :", time.Now().Weekday())

	if time.Now().Weekday() == time.Friday {
		fmt.Println("Yayy weekend")
	} else{
		fmt.Println("Damn")
	}

}

func varTry(){
	var c, python, java = true, false, "best"
	fmt.Println(c, python, java)
	//java = true
	//fmt.Println("java = ", java)
}
