package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Halo, Dunia!"

	// Mengubah string menjadi huruf kecil
	fmt.Println("Lowercase:", strings.ToLower(text))

	// Mengubah string menjadi huruf besar
	fmt.Println("Uppercase:", strings.ToUpper(text))

	// Mengecek apakah string dimulai dengan "Halo"
	fmt.Println("StartsWith 'Halo':", strings.HasPrefix(text, "Halo"))

	// Mengecek apakah string diakhiri dengan "Dunia!"
	fmt.Println("EndsWith 'Dunia!':", strings.HasSuffix(text, "Dunia!"))

	// Memisahkan string berdasarkan delimiter koma
	parts := strings.Split("apel,banana,cerry", ",")
	fmt.Println("Split by ',':", parts)

	// Mengganti bagian dari string
	newtext := strings.ReplaceAll(text, "Dunia", "Golang")
	fmt.Println("Replace:", newtext)
}
