package main

import "fmt"

func main() {

/*
	aList := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(aList); i++{
		aList[i] = aList[i] * aList[i]
	}
	fmt.Print(aList)
}
*/

	aList := []int{1, 2, 3, 4, 5, 6, 7}
	for i, _ := range aList{
		aList[i] = aList[i] * aList[i]
	}

	fmt.Print(aList)
}