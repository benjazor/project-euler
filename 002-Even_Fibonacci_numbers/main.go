package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n - 1) + fibo(n - 2)
	}
}

func main() {
	sum := 0
	for i:=0; fibo(i) < 4000000; i++{
		if fibo(i) % 2 == 0 {
			sum += fibo(i)
		}
	}
	fmt.Printf("%d", sum)
}