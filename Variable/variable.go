package main

import "fmt"

func main() {
	var name string = "Muzaki" // menampilkan tipe data string
	var umur int = 22          // menampilkan tipe data integer
	city := "Jakarta"          // menampilkan tipe data otomatis
	year := 2025               // menampilkan tipe data otomatis

	const pi = 3.14                 // Nilai Pi
	const appName string = "Golang" // Nama Aplikasi

	fmt.Println("Nama :", name)
	fmt.Println("Umur :", umur)
	fmt.Println("Kota :", city)
	fmt.Println("Tahun :", year)
	fmt.Println("Nilai Pi :", pi)
	fmt.Println("Aplikasi :", appName)
}
