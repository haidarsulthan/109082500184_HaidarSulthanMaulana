# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
<p align="center">[Haidar Sulthan Maulana] - [109082500184]</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
```go
package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
		}
	}

func permutasi(n, r int, hasil *int) {
	var nf, nrf	 int
	faktorial(n, &nf)
	faktorial(n-r, &nrf)
	*hasil =  nf / nrf
}

func kombinasi(n, r int, hasil *int) {
	var nf, rf, nrf int
	faktorial(n, &nf)
	faktorial(r, &rf)
	faktorial(n-r, &nrf)
	*hasil = nf / (rf * nrf)
}

func main() {
	var a, b, c, d, p1, c1, p2, c2 int
    fmt.Scan(&a, &b, &c, &d)

    permutasi(a, c, &p1)
    kombinasi(a, c, &c1)
    permutasi(b, d, &p2)
    kombinasi(b, d, &c2)

    fmt.Println(p1, c1)
    fmt.Println(p2, c2)
}

```
### Output Unguided :

##### Output 
![Output1](https://raw.githubusercontent.com/haidarsulthan/109082500184_HaidarSulthanMaulana/refs/heads/main/modul4/output/output1.png)
[penjelasan]
Program ini digunakan untuk menghitung permutasi dan kombinasi dari dua pasang bilangan yang diinputkan pengguna. Fungsi faktorial digunakan untuk menghitung n! dengan perulangan, dan menggunakan pointer agar hasil langsung disimpan ke variabel tanpa rturn(karna pointer berguna untuk menyimpan nilai sementara jadi tidak perlu mengembalikan nilai ke func lain). Fungsi permutasi dan kombinasi mengambil nilai dari func faktorial karna menggunakan pointer. Di fungsi main, program membaca empt bilangan, menghitung hasil untuk dua pasangan(permutasi dan kombinasi), lalu menampilkannya dalam dua baris.


'''
### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.
#### soal2.go

```go
package main

	import "fmt"

	func hitungSkor(soal *int, skor *int) {
		var time int
		*soal = 0
		*skor = 0
		for i := 0; i < 8; i++ {
			fmt.Scan(&time)
			if time <= 300 {
				*soal++
				*skor += time
			}
		}
	}
	
	func main() {
    var name, winner string
    var soal, skor, jmlsoal, jmlskor int
    jmlskor = 999999

	fmt.Scan(&name)

	for name != "selesai" {
		hitungSkor(&soal, &skor)
		
        if soal > jmlsoal || (soal == jmlsoal && skor < jmlskor) {
            jmlsoal = soal
            jmlskor = skor
            winner = name
    		}
		fmt.Scan(&name)
	}
		
	if winner != "" {
		fmt.Printf("%s %d %d\n", winner, jmlsoal, jmlskor)
		}
	}

```
### Output Unguided :

##### Output 
![Output1](https://raw.githubusercontent.com/haidarsulthan/109082500184_HaidarSulthanMaulana/refs/heads/main/modul4/output/output2.png)
[penjelasan]
program membaca nama peserta satu per satu sampai memasukkan "selesai" agar program berhenti. Untuk setiap peserta, program memanggil fungsi hitungSkor dengan i <= 8 karna ada 8 soal yang diberikan dan waktu <= 300 ini juga karna diberikan waktu 5 jam. Di fungsi hitungSkor, program membaca waktu dari 8 soal. Jika waktu ≤ 300, soal dihitung benar dan akan ditambahkan dengan waktu. Program ini memakai pointer supaya nilai soal dan skor bisa langsung berubah tanpa return. Setelah itu, program membandingkan hasil setiap peserta. Pemenang adalah yang soalnya paling banyak, dan pengerjaan yang cepat.
'''