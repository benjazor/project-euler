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
	number := 10001
	count := 0
	for i:=2; true; i++ {
		if prime(i) {
			count++
			
		}
		if count == number {
			fmt.Println(i)
			break
		}
	}
}