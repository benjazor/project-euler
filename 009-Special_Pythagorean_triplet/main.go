package main

import (
	"fmt"
	"math"
)

func main() {
	var number = 1000
	for a := 0; a < number/2; a++ {
		for b := 0; b < number/2; b++ {
			c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
			if c == math.Trunc(c) && a+b+int(c) == 1000 {
				fmt.Println(a * b * int(c))
				return
			}
		}
	}
}
