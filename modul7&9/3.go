package main

import "fmt"

func main() {
	var A, B string
	var skor1, skor2 int
	var hasil = []string{}
	fmt.Print("masukan nama club 1 : ")
	fmt.Scan(&A)
	fmt.Print("masukan nama club 2 : ")
	fmt.Scan(&B)

	for {

		fmt.Scan(&skor1, &skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}

		if skor1 > skor2 {
			hasil = append(hasil, "mu")

		} else if skor2 > skor1 {
			hasil = append(hasil, "inter")

		} else {
			hasil = append(hasil, "DRAW")

		}

	}
	for i := 0; i < len(hasil); i++ {
		fmt.Println("hasil", i+1, "=", hasil[i])
	}
}
