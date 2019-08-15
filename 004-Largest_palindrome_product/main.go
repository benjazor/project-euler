package main

import (
	"fmt"
	"strconv"
)

func palindromic(n int) bool{
	num := strconv.Itoa(n)
	for i:=0; i<len(num); i++{
		if num[i] != num[len(num)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	max := 0
	for i:=0; i<1000; i++{
		for j:=0; j<1000; j++{
			if palindromic(i*j) && i*j > max {
				max = i*j
			}
		}
	}
	fmt.Printf("%d\n", max)
}