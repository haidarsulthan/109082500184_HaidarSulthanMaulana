# <h1 align="center"> Laporan Praktikum Modul 10 </h1>
<p align="center">  [Haidar Sulthan Maulana] - [109082500184] </p>

## Unguided 

### 1. [KELINCI]
#### Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar.Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. 

```go
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



```
### Output Unguided :

##### Output 
![Output1](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul10/Output/Output1.png)

[Program Go tersebut digunakan untuk membaca data berat sejumlah anak kelinci, lalu menentukan berat terbesar dan terkecil dari data tersebut. Fungsi hitung menerima array berat kelinci dan jumlah data A, kemudian meminta input berat tiap anak kelinci. Setelah semua data dimasukkan, program membandingkan setiap elemen untuk mencari nilai maksimum dan minimum. Hasil akhirnya ditampilkan berupa berat kelinci terbesar dan terkecil. Fungsi main hanya bertugas meminta jumlah anak kelinci dari pengguna, mendeklarasikan array, lalu memanggil fungsi hitung untuk menjalankan proses perhitungan. Dengan demikian, program ini sederhana namun efektif untuk analisis data berat kelinci.]

### 2. [IKAN]
#### Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.


```go
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
```
### Output Unguided :

##### Output 

![Output2](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul10/Output/Output2.png)
[Program Go tersebut berfungsi untuk membagi data berat ikan ke dalam beberapa wadah dan menghitung rata-rata berat ikan di setiap wadah. Fungsi beratIkan pertama-tama meminta input jumlah ikan yang akan dijual (x) dan kapasitas wadah (y), kemudian membaca berat masing-masing ikan dan menyimpannya dalam array. Setelah itu, program melakukan iterasi dengan langkah sebesar kapasitas wadah (y), menghitung total berat ikan dalam wadah tersebut, lalu membagi dengan jumlah ikan di wadah untuk mendapatkan rata-rata. Hasilnya ditampilkan sebagai rata-rata berat ikan per wadah. Fungsi main hanya mendeklarasikan array ikan dan memanggil beratIkan.]

### 3. [BAYI]
#### Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. 

```go

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
```
### Output Unguided :

##### Output 

![Output3](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul10/Output/Output3.png)
[Program Go tersebut digunakan untuk mengolah data berat badan balita dengan cara membaca sejumlah input berat, lalu menghitung nilai minimum, maksimum, dan rata-rata dari data tersebut. Fungsi hitungMinMax menerima array berat balita dan pointer untuk menyimpan nilai minimum, maksimum, serta jumlah balita; fungsi ini meminta input jumlah data, mengisi array dengan berat balita, kemudian menentukan berat terkecil dan terbesar. Fungsi rerata menghitung rata-rata dengan menjumlahkan semua elemen array lalu membaginya dengan jumlah balita. Di dalam main, program memanggil hitungMinMax untuk menampilkan berat minimum dan maksimum, kemudian memanggil rerata untuk menampilkan rata-rata berat balita. Dengan demikian, program ini membantu menganalisis data berat balita secara sederhana.]
