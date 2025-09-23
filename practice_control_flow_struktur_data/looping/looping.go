package main

import "fmt"

func main() {
	// FOR LOOP
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	// WHILE LOOP (menggunakan for dengan kondisi)
	jumlah := 1
	for jumlah <= 5 {
		fmt.Println("Jumlah saat ini:", jumlah)
		jumlah++
	}

	// RANGE (untuk slice/array/map)
	nama := []string{"Alice", "Bob", "Charlie"}
	for index, value := range nama {
		fmt.Printf("Index: %d, Nama: %s\n", index, value)
	}
}
