# <h1 align="center"> Laporan Praktikum Modul 10 </h1>
<p align="center">  [Haidar Sulthan Maulana] - [109082500184] </p>

## Unguided 

### 1. [INTEGER]
#### Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.. 

```go
package main

import "fmt"

const NMAX = 1000

type tabint [NMAX]int

func insertionSort(arr *tabint, n int) {
	for i := 1; i < n; i++ {
		key := (*arr)[i]
		j := i - 1
		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

func cekJarak(arr tabint, n int) {
	if n > 1 {
		diff := arr[1] - arr[0]
		status := 1

		for i := 1; i < n-1; i++ {
			if arr[i+1]-arr[i] != diff {
				status = 0
			}
		}

		if status == 1 {
			fmt.Printf("Data berjarak %d\n", diff)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

func main() {
	var arr tabint
	var n int
	var num int

	fmt.Scan(&num)
	for num >= 0 {
		arr[n] = num
		n++
		fmt.Scan(&num)
	}

	insertionSort(&arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	cekJarak(arr, n)
}
```
### Output Unguided :

##### Output 
![Output1](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul14-2/Output/Output1.png)

[Program meminta input angka acak dan akan mulai mencetak ketika input < 0. Kemudian program mengurutkan output menggunakan metode sorting insertion sort(Ascending yaitu terurut membesar). temp berfungsi sebagai penyimpan sementara elemen yang sedang diproses, supaya nilainya tidak hilang ketika elemen-elemen lain digeser. Kemudian akan di cetak outputnya terurut membesar dan juga akan dihitung jarak per-angkanya misal 1 7 13 19 25 31 37 43 urutan angka ini memili jarak 6.]

### 2. [BUKU]
#### Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan.


```go
package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax + 1]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i <= n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n > 0 {
		maxRating := pustaka[1].rating
		for i := 2; i <= n; i++ {
			if pustaka[i].rating > maxRating {
				maxRating = pustaka[i].rating
			}
		}

		for i := 1; i <= n; i++ {
			if pustaka[i].rating == maxRating {
				fmt.Printf("%s, %s, %s, %d\n", pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun)
			}
		}
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 2; i <= n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 1 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 1; i <= limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 1
	kanan := n
	found := -1

	for kiri <= kanan && found == -1 {
		med := (kiri + kanan) / 2
		if pustaka[med].rating == r {
			found = med
		} else if pustaka[med].rating < r {
			kanan = med - 1
		} else {
			kiri = med + 1
		}
	}

	if found != -1 {
		fmt.Printf("%s, %s, %s, %d, %d, %d\n", pustaka[found].judul, pustaka[found].penulis, pustaka[found].penerbit, pustaka[found].tahun, pustaka[found].eksemplar, pustaka[found].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var n, rating int
	var pustaka DaftarBuku

	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	fmt.Scan(&rating)

	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, rating)
}

```
### Output Unguided :

##### Output 

![Output3](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul11-2/Output/Output2.png)
[Program golang diatas berfungsi untuk mengelola data buku dengan beberapa fitur utama. Pertama, tipe data Buku menyimpan informasi lengkap seperti id, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Data buku disimpan dalam array DaftarBuku dengan kapasitas maksimum 7919. Fungsi DaftarkanBuku digunakan untuk membaca input data buku dari pengguna sesuai jumlah yang dimasukkan. Selanjutnya, CetakTerfavorit mencari buku dengan rating tertinggi dan mencetak detailnya. Fungsi UrutBuku mengurutkan daftar buku berdasarkan rating secara menurun menggunakan algoritma insertion sort. Setelah diurutkan, Cetak5Terbaru menampilkan lima judul buku teratas, atau kurang jika jumlah buku tidak mencapai lima. Fungsi CariBuku melakukan pencarian dengan metode binary search untuk menemukan buku dengan rating tertentu, lalu menampilkan detailnya jika ditemukan. Terakhir, fungsi main mengatur alur program dengan membaca jumlah buku, mendaftarkan data, membaca rating target, lalu menjalankan semua fungsi tersebut secara berurutan.]
