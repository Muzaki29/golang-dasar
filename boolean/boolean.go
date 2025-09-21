package main

import "fmt"

func main() {
	var isStudent bool = true    // boolean true
	var isGraduated bool = false // boolean false

	fmt.Println("Apakah mahasiswa?", isStudent)
	fmt.Println("Apakah sudah lulus?", isGraduated)

	// contoh boolean dari perbandingan
	var umur int = 20
	var isAdult bool = umur >= 18 // hasilnya true karena 20 >= 18

	fmt.Println("Apakah dewasa?", isAdult)
}
