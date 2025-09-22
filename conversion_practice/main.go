package main

import (
	"fmt"
	"strconv"
)

func main() {
	// =============================
	// 1. Int ↔ String
	// =============================
	angka := 123
	strAngka := strconv.Itoa(angka) // int → string
	fmt.Println("Int ke String :", strAngka)

	angkaBaru, err := strconv.Atoi("456") // string → int
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("String ke Int :", angkaBaru)
	}

	// =============================
	// 2. Bool ↔ String
	// =============================
	truth := true
	strBool := strconv.FormatBool(truth) // bool → string
	fmt.Println("Bool ke String :", strBool)

	valBool, err := strconv.ParseBool("true") // string → bool
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("String ke Bool :", valBool)
	}

	// =============================
	// 3. Float ↔ String
	// =============================
	pi := 3.14159
	strFloat := strconv.FormatFloat(pi, 'f', 2, 64) // float64 → string (2 decimal)
	fmt.Println("Float ke String :", strFloat)

	valFloat, err := strconv.ParseFloat("2.718", 64) // string → float64
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("String ke Float :", valFloat)
	}
}
