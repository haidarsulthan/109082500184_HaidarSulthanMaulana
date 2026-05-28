# <h1 align="center"> Laporan Praktikum Modul 10 </h1>
<p align="center">  [Haidar Sulthan Maulana] - [109082500184] </p>

## Unguided 

### 1. [HERCULES]
#### Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. 

```go
package main

import "fmt"

type arrInt [1000000]int

func selectionAsc(T *arrInt, n int) {
	var idxMin int

	i := 1
	for i <= n-1 {
		idxMin = i - 1
		j := i
		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j++
		}
		t := T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = t
		i++
	}
}

func main() {
	var n, m int
	var a arrInt

	fmt.Scan(&n)

	i := 0
	for i < n {
		fmt.Scan(&m)

		j := 0
		for j < m {
			fmt.Scan(&a[j])
			j++
		}

		selectionAsc(&a, m)

		j = 0
		for j < m {
			fmt.Print(a[j], " ")
			j++
		}
		fmt.Println()

		i++
	}

}

```
### Output Unguided :

##### Output 
![Output1](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul14/Output/Output1.png)

[Program menginputkan n, lalu melakukan perulangan sebanyak n kali. Di setiap iterasi, dibaca integer m (jumlah rumah kerabat), diikuti input m angka yang disimpan ke array a. Lalu, fungsi selectionAsc dipanggil. Variabel i dimulai dari 1, perulangan menggunakan i<=n-1, kemudian nilai awal minimum(idxMin)=i-1, masuk ke perulangan j dimulai dari i, di dalam perulangan j nilai dibandingan untuk mendapatkan nilai terkecil ke terbesar nilai terus beftambah sampai n habis dikurangi 1 dalam perulangan i.]

### 2. [HERCULES TAKUT MENYEBRANG]
#### Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.


```go
package main

import "fmt"

type arrInt [100000]int

func selectionAsc(T *arrInt, n int) {
	var idxMin int

	i := 1
	for i <= n-1 {
		idxMin = i - 1
		j := i
		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j++
		}
		t := T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = t
		i++
	}
}
func main() {
	var n, m int
	var a arrInt

	fmt.Scan(&n)

	i := 0
	for i < n {
		fmt.Scan(&m)

		j := 0
		for j < m {
			fmt.Scan(&a[j])
			j++
		}

		selectionAsc(&a, m)

		j = 0
		for j < m {
			if a[j]%2 != 0 {
				fmt.Print(a[j], " ")
			}

			j++
		}
		j = m - 1

		for j >= 0 {

			if a[j]%2 == 0 {
				fmt.Print(a[j], " ")
			}

			j--
		}

		fmt.Println()

		i++
	}
}

```
### Output Unguided :

##### Output 

![Output2](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul14/Output/Output2.png)
[Input sama seperti Soal 1. Yang berbeda adalah cara mencetak hasilnya. Pada func selectionAsc samapersis juga dengan soal 1. Yang membuat beda adalah soal 2 memiliki keluaran membesar ganjil dan mengecil genap.
Ganjil Ascending: Loop dimulai dari menentukan nilai (j = 0) selanjutnya (j < m). jika a[j] % 2 != 0 (ganjil), langsung dicetak. Karena array sudah ascending, bilangan ganjil yang keluar otomatis terurut membesar.
Genap Descending: Loop dimulai dari menentukan (j = m-1) selanjutnya (j >= 0). jika a[j] % 2 == 0 (genap), langsung dicetak. Karena array sudah ascending dan kita baca dari kanan, bilangan genap yang keluar otomatis terurut mengecil.]

### 3. [KOMPETISI]
#### Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya? 

```go

package main

import "fmt"

type arr [1000000]int

func insertionSort(T *arr, n int) {
	i := 1
	for i <= n-1 {
		j := i
		temp := T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func Median(T *arr, n int) int {
	var median int
	if n%2 != 0 {
		median = T[n/2]
	} else {
		median = (T[(n/2)-1] + T[n/2]) / 2
	}
	return median
}

func main() {
	var data arr
	var angka, n int

	fmt.Scan(&angka)

	for angka != -5313 {
		if angka == 0 {

			insertionSort(&data, n)

			fmt.Println(Median(&data, n))
		} else {

			data[n] = angka
			n = n + 1
		}

		fmt.Scan(&angka)
	}
}

```
### Output Unguided :

##### Output 

![Output3](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul14/Output/Output3.png)
[Program digunakan untuk menyimpan sekumpulan angka, mengurutkannya dengan algoritma insertion sort, lalu menampilkan nilai median setiap kali input 0 diberikan. Program dimulai dengan mendefinisikan tipe data arr sebagai array berukuran besar ([1000000]int) agar bisa menampung banyak angka. Fungsi insertionSort bertugas mengurutkan isi array dari kecil ke besar menggunakan metode insertion sort, yaitu dengan mengambil satu elemen (temp) lalu menggesernya ke posisi yang benar di bagian array yang sudah terurut. Fungsi Median digunakan untuk mencari nilai tengah data: jika jumlah data n ganjil maka median adalah elemen tepat di tengah (T[n/2]), sedangkan jika genap maka median adalah rata-rata dua elemen tengah. Pada fungsi main, program terus membaca input angka menggunakan fmt.Scan. Jika angka yang dimasukkan bukan 0, angka tersebut disimpan ke array data dan jumlah data n ditambah satu. Jika input adalah 0, program akan mengurutkan data yang sudah masuk lalu mencetak mediannya. Proses ini terus berlangsung sampai pengguna memasukkan angka -5313 sebagai tanda berhenti.]
