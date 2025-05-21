package main

import (
  "fmt"
  "math/rand"
)

func massive() []int {
  var length int
  fmt.Print("Введите количество элементов в массиве: ")
  fmt.Scan(&length)

  list := make([]int, length)

  for i, _ := range list {
    list[i] = rand.Intn(51)
  }
  fmt.Println("Список до изменений -", list)
  return list
}

func main() {
  var lst = massive()

  var f func(list []int) []int = new_massive
  fmt.Println(f(lst))
}

func new_massive(list []int) []int {
  for i, nums := range list {
    if nums == 50 {
      list[i] = 200
    }
  }
  return list
}
