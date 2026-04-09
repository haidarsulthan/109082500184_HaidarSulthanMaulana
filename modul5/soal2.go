package main

import "fmt"

func sangbintang(n int) {
	if n == 0 {
		return
	}
	sangbintang(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
func main() {
	var n int
	fmt.Scan(&n)
	sangbintang(n)
}
