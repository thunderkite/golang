package main

import "fmt"

func main(){
	fmt.Println(DomainForLocale("site.com", ""))
}


func DomainForLocale(domain, locale string) string{
	if locale == ""{
		locale = "en"
		return locale + "." + domain
	}
	return locale + "." + domain
}