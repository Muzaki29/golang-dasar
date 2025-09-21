package main

import "fmt"

func main() {
	var nama string = "Muzaki"
	var hobi string = "Lari"

	fmt.Println("Nama:", nama)
	fmt.Println("Hobi:", hobi)

	// gabung string
	pesan := "Halo, " + nama + "! Selamat belajar "
	fmt.Println(pesan)

	// panjang string
	fmt.Println("Panjang nama:", len(nama))

	// ambil karakter pertama (dalam bentuk byte)
	fmt.Println("Karakter pertama:", string(nama[0]))
}
