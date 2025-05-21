package main

import "fmt"

func main() {
	var N, X int
	// Считываем значения N и X
	fmt.Scan(&N, &X)

	// Создаем срез длиной 0, но с емкостью N для оптимизации
	result := make([]int, 0, N)

	// Заполняем срез числом X N раз
	for i := 0; i < N; i++ {
		result = append(result, X)
	}

	// Выводим результат в формате [X,X,X...]
	// Сначала выводим открывающую скобку
	fmt.Print("[")
	// Выводим все элементы через запятую
	for i, num := range result {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(num)
	}
	// Закрываем скобку
	fmt.Println("]")
}
