package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func main(){
	fmt.Println(Greetings("жЕнЯ"))
}

func Greetings(name string) string{
	processedName := cases.Title(language.Russian).String(strings.ToLower(strings.Trim(name, " ")))
	return "Привет, " + processedName + "!"
}