# <h1 align="center"> Laporan Praktikum Modul 9 </h1>
<p align="center">  [Haidar Sulthan Maulana] - [109082500184] </p>

## Unguided 

### 1. [PUSAT_LINGKARAN]
#### Suatu lingkaran didefinisikan dengan koordinat titik pusat (𝑐𝑥,𝑐𝑦) dengan radius 𝑟. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (𝑥,𝑦) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya. 

```go
  package main

import "fmt"

func dalamLingkaran(cx, cy, r, x, y int) bool {
	dx := x - cx
	dy := y - cy
	return dx*dx+dy*dy <= r*r
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := dalamLingkaran(cx1, cy1, r1, x, y)
	dalam2 := dalamLingkaran(cx2, cy2, r2, x, y)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}



```
### Output Unguided :

##### Output 
![Output1](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul10/Output/Output1.png)

[Program tersebut bekerja dengan cara membaca input pusat dan radius dua lingkaran serta koordinat sebuah titik, lalu menggunakan fungsi dalamLingkaran untuk menghitung apakah titik tersebut berada di dalam lingkaran. hasilnya berupa boolean yang menunjukkan apakah titik ada di lingkaran pertama, kedua, keduanya, atau tidak sama sekali, kemudian bagian main mencetak pesan sesuai kondisi: “Titik di dalam lingkaran 1 dan 2”, “Titik di dalam lingkaran 1”, “Titik di dalam lingkaran 2”, atau “Titik di luar lingkaran 1 dan 2”.]

### 2. [BILANGAN]
#### Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu.


```go
	package main

import "fmt"

func e(arr []int, idx int) []int {
	n := len(arr)

	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]

	}
	arr = arr[:n-1]
	return arr

}
func avg(arr []int) float64 {
	n := len(arr)
	jumlah := 0

	for i := 0; i < n; i++ {
		jumlah += arr[i]

	}

	return float64(jumlah) / float64(n)

}

func freq(arr []int) ([]int, []int) {
	n := len(arr)
	nilai := []int{}
	hit := []int{}

	for i := 0; i < n; i++ {
		sudah := false
		for j := 0; j < i; j++ {

			if arr[i] == arr[j] {
				sudah = true

				break
			}
		}

		if sudah {
			continue
		}

		hitung := 0

		for j := 0; j < n; j++ {
			if arr[i] == arr[j] {
				hitung++
			}
		}
		nilai = append(nilai, arr[i])
		hit = append(hit, hitung)

	}
	return nilai, hit
}

func main() {

	var x, index int

	fmt.Print("masukin angka index yng mau di hapus : ")

	fmt.Scan(&index)

	fmt.Print("masukin angka kelipatan : ")

	fmt.Scan(&x)

	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)

	fmt.Println("ini jawban soal a")

	for i := 0; i < n; i++ {
		fmt.Printf("index %d dan valuenya %d\n", i, arr[i])

	}
	fmt.Println()

	fmt.Println("ini jawban soal b dan c")
	fmt.Println()

	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Printf("menampilkan value index ganjil %d\n", arr[i])
		} else {
			fmt.Printf("menampilkan value index genap %d\n", arr[i])

		}

	}
	fmt.Println()

	fmt.Println("ini jawban soal d")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("menampilkan value index ganjil %d\n", arr[i])
		}
	}
	fmt.Println()

	fmt.Println("ini jawban soal e")
	fmt.Println("menghapus index")

	arr = e(arr, index)

	fmt.Println(arr)
	fmt.Println()

	fmt.Println("ini jawban soal f")
	fmt.Println("avg atau rata rata index a-z")

	avg := avg(arr)
	fmt.Println("rata ratanya : ", avg)

	vals, counts := freq(arr)
	for i := 0; i < len(vals); i++ {
		fmt.Println(vals[i], "muncul", counts[i])
	}

}


```
### Output Unguided :

##### Output 

![Output2](https://github.com/haidarsulthan/109082500184_HaidarSulthanMaulana/blob/main/modul10/Output/Output2.png)
[Program Go ini melakukan beberapa operasi pada slice arr: pertama menampilkan setiap index dan nilainya, lalu memisahkan tampilan berdasarkan index genap/ganjil serta index kelipatan x, kemudian menghapus elemen pada index tertentu menggunakan fungsi e (dengan cara menggeser elemen ke kiri dan memotong slice), setelah itu menghitung rata-rata dengan fungsi avg, dan terakhir menghitung frekuensi kemunculan tiap nilai unik memakai fungsi freq; meskipun secara konsep sudah benar untuk latihan dasar, kode ini masih punya kelemahan seperti tidak ada validasi input (bisa error kalau index tidak valid atau x = 0), penggunaan algoritma yang kurang efisien pada freq (O(n²)), serta penamaan yang membingungkan seperti variabel avg yang menimpa nama fungsi.]

### 3. [PERTANDINGAN_BOLA]
#### Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. 

```go

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



```
### Output Unguided :

##### Output 

![Output1](https://raw.githubusercontent.com/haidarsulthan/109082500184_HaidarSulthanMaulana/refs/heads/main/modul7%269/Output/Output3.png)
[Program ini membaca dua nama klub (A dan B), lalu masuk ke loop tanpa batas untuk menerima input skor pertandingan berulang (skor1 dan skor2); setiap pasangan skor dibandingkan—jika skor1 > skor2 maka hasil "mu" ditambahkan ke slice hasil, jika skor2 > skor1 maka "inter", dan jika sama maka "DRAW"—proses ini berhenti saat salah satu skor bernilai negatif, lalu seluruh hasil pertandingan ditampilkan satu per satu; secara logika dasar sudah jalan, tapi ada asumsi keliru yang cukup serius: pemenang tidak menggunakan variabel A dan B melainkan hardcoded "mu" dan "inter", jadi input nama klub sebenarnya tidak dipakai sama sekali, selain itu tidak ada validasi input (misalnya jika user salah format input), dan loop for {} tanpa kontrol tambahan membuat program sepenuhnya bergantung pada input negatif untuk berhenti.]


### 4. [REVERSE_STRING]
#### Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom. 

```go
 package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for {
		fmt.Scanf("%c", &ch)

		if ch == '.' {
			break
		}

		if ch == '\n' || ch == ' ' {
			continue
		}

		(*t)[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Teks: ")
	cetakArray(tab, m)

	if palindrom(tab, m) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}

	balikanArray(&tab, m)

	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)
}



```
### Output Unguided :

##### Output 
![Output_4](https://raw.githubusercontent.com/haidarsulthan/109082500184_HaidarSulthanMaulana/refs/heads/main/modul7%269/Output/Output4.png)

[Program ini membaca karakter satu per satu ke dalam array tabel (maksimal 127 rune) lewat isiArray, berhenti saat menemui titik (.) dan mengabaikan spasi serta newline; jumlah karakter disimpan di m, lalu cetakArray menampilkan isi array, fungsi palindrom mengecek apakah teks sama jika dibalik dengan membandingkan elemen dari depan dan belakang, dan balikanArray benar-benar membalik isi array dengan menukar elemen simetris; secara logika ini sudah benar untuk manipulasi array dan pengecekan palindrom, tetapi ada asumsi yang perlu dikritisi: input sangat bergantung pada format per karakter dengan Scanf("%c") yang rawan error (misalnya buffering newline), tidak ada batasan jika input melebihi NMAX (bisa overflow indeks), serta penghapusan spasi membuat hasil palindrom berbeda dari ekspektasi jika pengguna menganggap spasi bagian dari teks (misalnya “katak katak” diperlakukan berbeda).]

