package main

import "fmt"

func main() {
	var a, b, c, d string
	berhasil := true
	for i := 1; i <= 5; i++ {
		fmt.Print("percobaan ", i, ": ")
		fmt.Scan(&a, &b, &c, &d)
		if !(a == "merah" && b == "kuning" && c == "hijau" && d == "ungu") {
			berhasil = false
		}
	}
	if berhasil {
		fmt.Print("BERHASIL : true")
	} else {
		fmt.Print("BERHASIL : false")
	}
}
