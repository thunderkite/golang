package main

import (
	"fmt"
	"errors"
	
)

func main() {
	result, err := divide()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}

func divide() (int, error) {
	x, y := write_terminal()
	if x == 0 || y == 0 {
		return 0, errors.New("вы не можете делить на 0!")
	}
	return x / y, nil
}

func write_terminal() (int, int) {
	var x, y int
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&x)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&y)
	return x, y
}

