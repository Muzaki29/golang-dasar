package main

import (
	"fmt"
)

func main() {
	//signed integer
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = 100 // int sama dengan int32 atau int64 tergantung dari sistem operasi

	//unsigned integer
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	var ui uint = 100 // uint sama dengan uint32 atau uint64 tergantung dari sistem operasi

	//Menampilkan nilai
	fmt.Printf("signed integers:")
	fmt.Printf("int8 	: %v\n", i8)
	fmt.Printf("int16	: %v\n", i16)
	fmt.Printf("int32	: %v\n", i32)
	fmt.Printf("int64	: %v\n", i64)
	fmt.Printf("int		: %v\n", i)

	fmt.Println("\nunsigned integers:")
	fmt.Printf("uint8	: %v\n", ui8)
	fmt.Printf("uint16	: %v\n", ui16)
	fmt.Printf("uint32	: %v\n", ui32)
	fmt.Printf("uint64	: %v\n", ui64)
	fmt.Printf("uint	: %v\n", ui)
}
