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
