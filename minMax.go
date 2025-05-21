package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Наименьшее число: %.2f\n", MinInt(4, 2))
}

func MinInt(x, y float64) float64{
	minNum := math.Min(x, y)
	return minNum
}