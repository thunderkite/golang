package main

import "fmt"

func main() {
    numbers := []int{5, 10, 15, 20, 25, 50, 20}
    
    // Итерируемся по срезу, используя range для лучшей читаемости
    for i := range numbers { // num = numbers[i]!!!
        // Проверяем, равен ли текущий элемент 50
        if numbers[i] == 50 {
            // Заменяем значение в исходном срезе
            numbers[i] = 200
        }
    }
    
    // Выводим измененный срез
    fmt.Println("Измененный список:", numbers)
}