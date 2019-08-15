package main

import "fmt"

func evenlyDivisible(n int, x int) bool{
	for i:=1; i<x; i++ {
		if n % i != 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int = 20
	for i:=1; true; i++ {
		if evenlyDivisible(i, number) {
			fmt.Printf("%d", i)
			break
		}
	}
}