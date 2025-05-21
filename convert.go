package main

import( 
	"strconv"
	"fmt"
	"reflect"
)



func main(){
	fmt.Println(reflect.TypeOf(intToString(-42)))
}

func intToString(number int) string{
	string := strconv.Itoa(number)
	return string
}