package main

import "fmt"


func main() {
	var N int
	fmt.Scan(&N)
	var lst []int
	for i := 1; i < N + 1; i++ {
		lst = append(lst, i*i)
	}
	fmt.Println(lst)
}

/*
func main() {
	var N int
	fmt.Scan(&N)

	
	// Создаем срез для хранения квадратов
	squares := make([]int, 0, N)
	
	// Заполняем срез квадратами чисел от 1 до N включительно
	for i := 1; i <= N; i++ {
		squares = append(squares, i*i)
	}
	
	// Выводим результат в требуемом формате
	fmt.Print("[")
	for i, num := range squares {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(num)
	}
	fmt.Println("]")
}
*/