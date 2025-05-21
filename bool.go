package main

import "fmt"

func main(){
	fmt.Println(IsValid(1, "hello"))
}


func IsValid(id int, text string) bool {
	if id <= 0 || text == ""{
		return false
	}
	return true
}