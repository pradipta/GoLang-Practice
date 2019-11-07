package main

import (
	"crypto/tls"
	"fmt"
)
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
}

func mapTry(mode string){
	modeMap := map[string]tls.RenegotiationSupport{
		"once":   tls.RenegotiateOnceAsClient,
		"freely": tls.RenegotiateFreelyAsClient,
		"never":  tls.RenegotiateNever,
	}
	if val, ok := modeMap[mode]; ok {
		if val == val {

		}
		fmt.Println("found")
	} else {
		//if anything other than the allowed values is passed, it'll thrown an error
		fmt.Println("not found")
	}
}

func mapTryNoElse(mode string){
	modeMap := map[string]tls.RenegotiationSupport{
		"once":   tls.RenegotiateOnceAsClient,
		"freely": tls.RenegotiateFreelyAsClient,
		"never":  tls.RenegotiateNever,
	}


	val, ok := modeMap[mode]
	if val == val{

	}
	if !ok{
		if val == val{

		}
		fmt.Println("herenotfound")
		return
	}
	fmt.Println("found")

	fmt.Println(ok)
}
