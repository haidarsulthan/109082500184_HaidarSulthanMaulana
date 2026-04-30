package main

import "fmt"

type arrBalita [100]float64
func hitungMinMax(arrBerat *arrBalita, bMin, bMax *float64, balita *int) {
	
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(balita)
	for i := 0; i < *balita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ",i+1)
		fmt.Scan(&arrBerat[i])
	}
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < *balita; i++ {
		if arrBerat[i] < *bMin {	
			*bMin = arrBerat[i]
		}else if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
		
	}
		fmt.Printf("Berat balita minimum: %.2f", *bMin)
		fmt.Println()
		fmt.Printf("Berat balita maksimum: %.2f", *bMax)
		fmt.Println()
}
	
func rerata(arrBerat *arrBalita, balita int) float64 {
	var total float64

	if balita == 0 {
		return 0
	}
	for i := 0; i < balita; i++ {
		total = total + arrBerat[i]
	}
	return total / float64(balita)
	
}
func main() {
	var arrBerat arrBalita
	var bMin, bMax float64
	var balita int
	hitungMinMax(&arrBerat, &bMin, &bMax, &balita)
	fmt.Printf("Rata-rata berat balita: %.2f", rerata(&arrBerat, balita))
}