package main

import "fmt"
import "time"

func main(){
	fmt.Println("Time is : ", time.Now())
	fmt.Println("Time2 is :", time.Now().Weekday())

	if time.Now().Weekday() == time.Friday {
		fmt.Println("Yayy weekend")
	} else{
		fmt.Println("Damn")
	}

}
