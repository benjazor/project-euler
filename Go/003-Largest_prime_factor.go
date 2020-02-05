package main

import "fmt"

func prime(n int) bool {
	for i:=2; i<n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	number := 600851475143
	for i:=1; i <= number; i++ {
		if prime(i) && number % i == 0{
			number /= i
			fmt.Printf("%d\n", i)
		}
	}
}