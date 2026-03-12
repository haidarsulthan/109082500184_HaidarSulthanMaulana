package main

import "fmt"

func main() {
	var parsel, biaya, detail, total int

	fmt.Print("Berat parsel dalam gram: ")
	fmt.Scan(&parsel)

	kg := parsel / 1000
	gr := parsel % (kg * 1000)
	biaya = kg * 10000

	fmt.Println("Detail berat:", kg, "kg", "+", gr, "gr")

	if kg > 10 {
		detail = gr * 5
		total = biaya
	} else if gr >= 500 {
		detail = gr * 5
		total = biaya + detail
	} else {
		detail = gr * 15
		total = biaya + detail
	}

	fmt.Println("detail biaya: ", "Rp.", biaya, " + Rp.", detail)
	fmt.Print("Total biaya: Rp.", total)
}
