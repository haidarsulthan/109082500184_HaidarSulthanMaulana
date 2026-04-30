package main

import "fmt"

func beratIkan(ikan [1000]float64) {
	var x, y int
	fmt.Print("Ikan yang akan dijual (x): ")
	fmt.Scan(&x)
	fmt.Print("Wadah (y): ")
	fmt.Scan(&y)

		fmt.Println()

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}
	fmt.Println("Pembagian wadah: ")
	for i := 0; i < x; i += y {
		var total float64 = 0
		var jumlah int = 0

		for j := i; j < i+y && j < x; j++ {
			total += ikan[j]
			jumlah++
		}
		fmt.Printf("Wadah %d: %.2f kg\n", (i/y)+1, total/float64(jumlah))
	}
}
func main() {
	var ikan [1000]float64
	beratIkan(ikan)
}