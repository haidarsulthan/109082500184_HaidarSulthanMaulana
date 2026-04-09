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