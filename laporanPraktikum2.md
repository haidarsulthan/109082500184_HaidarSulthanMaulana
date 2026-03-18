# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Haidar Sulthan Maulana] - [109082500184]</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
```go
package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutasi(a, c), kombinasi(a, c))
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}

```
### Output Unguided :

##### Output 
![Output1](https://raw.githubusercontent.com/haidarsulthan/109082500184_HaidarSulthanMaulana/main/Modul2/output/Output1.png)
[penjelasan]
Program ini meminta 3 output dari user lalu inputan nya diurutkan sesuai masukkan user dan akhirnya memberi outputan berupa hasil geser variabel.


'''
### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x^2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function.
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	hasil1 := f(g(h(a)))
	hasil2 := g(h(f(b)))
	hasil3 := h(f(g(c)))

	fmt.Println(hasil1)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}

```
### Output Unguided :

##### Output 
![Output1](https://raw.githubusercontent.com/haidarsulthan/109082500184_HaidarSulthanMaulana/refs/heads/main/Modul2/output/Output2.png)
[penjelasan]
Program ini dibuat untuk mengecek kombinasi 4 warna dalam percobaan. Jika warnanya merah, kuning, hijau, ungu maka true. Sebaliknya jika tidak sesuai maka false.


'''
### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut.
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}
func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)

	in1 := didalam(cx1, cy1, r1, x, y)
	in2 := didalam(cx2, cy2, r2, x, y)

	if in1 && in2 {
		fmt.Println("Titik didalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik didalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik didalam lingkaran 2")
	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
![Output1](https://raw.githubusercontent.com/haidarsulthan/109082500184_HaidarSulthanMaulana/refs/heads/main/Modul2/output/Output3.png)
[penjelasan]
Program ini dibuat untuk menghitung biaya pengiriman berdasarkan berat dalam gram. Program mengubah gram menjadi kilogram lalu mengalikan dengan harga perkilonya lalu ditambahkan sisa beratnya dan program menampilkan hasil jumlahnya menjadi total harga.