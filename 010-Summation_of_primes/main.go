package main

import "fmt"

func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	result := 0

	for i := 2; i < 2000000; i++ {
		if prime(i) {
			result += i
		}
	}

	fmt.Println(result)
}
