package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var N int
	fmt.Scan(&N)
	
	// Создаем два среза:
	// - один для хранения изначальных значений
	// - другой для хранения результата
	initially := make([]int, 0, N)
	result := make([]int, 0, N)
	
	// Заполняем срез изначальными (рандомными) значениями от 1 до N
	for i := 1; i < N + 1; i++ {
		initially = append(initially, rand.Intn(100))
	}
	
	// Итерируемся по изначальному срезу, если число четное - его квадрат
	// добавляем в результат
	for i, num := range initially {
		if num % 2 == 0 {
			result = append(result, num * num)
		}
		initially[i] = num
	}
	
	// Выводим результат
	fmt.Println(result)
}