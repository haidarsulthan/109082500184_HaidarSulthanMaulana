package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
		}
	}

func permutasi(n, r int, hasil *int) {
	var nf, nrf	 int
	faktorial(n, &nf)
	faktorial(n-r, &nrf)
	*hasil =  nf / nrf
}

func kombinasi(n, r int, hasil *int) {
	var nf, rf, nrf int
	faktorial(n, &nf)
	faktorial(r, &rf)
	faktorial(n-r, &nrf)
	*hasil = nf / (rf * nrf)
}

func main() {
	var a, b, c, d, p1, c1, p2, c2 int
    fmt.Scan(&a, &b, &c, &d)

    permutasi(a, c, &p1)
    kombinasi(a, c, &c1)
    permutasi(b, d, &p2)
    kombinasi(b, d, &c2)

    fmt.Println(p1, c1)
    fmt.Println(p2, c2)
}