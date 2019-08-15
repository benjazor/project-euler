package main

import (
	"fmt"
	"math"
)

func sumOfTheSquares(n int) int{
	result := 0
	for i:=1; i<=n; i++ {
		result += int(math.Pow(float64(i),2))
	}
	return result
}

func squareOfTheSum(n int) int{
	result := 0
	for i:=1; i<=n; i++ {
		result += i
	}
	return int(math.Pow(float64(result), 2))
}

func main() {
	var number int = 100
	fmt.Printf("%f\n", math.Abs(float64(sumOfTheSquares(number)-squareOfTheSum(number))))
}