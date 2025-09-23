package main

import "fmt"

func main() {
	// ARRAY (fixed length)
	var angka [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", angka)

	// SLICE (dynamic length)
	buah := []string{"apel", "pisang", "jeruk"}
	buah = append(buah, "mangga") // menambahkan elemen baru
	fmt.Println("Slice:", buah)

	// MAP (key-value pairs)
	mahasiswa := map[string]int{
		"Qois":     21,
		"Naykilla": 22,
	}
	mahasiswa["Charlie"] = 23 // menambahkan elemen baru
	fmt.Println("Map:", mahasiswa)
}
