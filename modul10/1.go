package main

import "fmt"

func hitung(kelinci[1000]float64, A int) {
	for i := 0; i < A; i++{
		fmt.Printf("Anak kelinci ke %d: ",i+1)
		fmt.Scan(&kelinci[i])
	}
	max := kelinci[0]
	min := kelinci[0]
	for i := 1; i < A; i++{
	if kelinci[i] > max {
		max = kelinci[i]
	}else if kelinci[i] < min {
		min = kelinci[i]
	}
}
	fmt.Println("Berat kelinci terbesar: ",max)
	fmt.Print("Berat kelinci terkecil: ",min)

}

func main() {
	var kelinci[1000] float64
	var N int
	fmt.Print("Jumlah anak kelinci: ")
	fmt.Scan(&N)

	hitung(kelinci, N)

}