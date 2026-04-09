package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Print("n: ")
	for i := 0; i <= n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Print("Sn: ")
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}

}
