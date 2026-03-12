# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Haidar Sulthan Maulana] - [109082500184]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main
import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/Modul2/output/Output1.png)
[penjelasan]
Program ini meminta 3 output dari user lalu inputan nya diurutkan sesuai masukkan user dan akhirnya memberi outputan berupa hasil geser variabel.