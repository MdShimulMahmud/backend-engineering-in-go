// datatypes.go
package main

import "fmt"

func main() {
	// Integer types
	var a int = 10
	var b int8 = 20
	var c int16 = 30
	var d int32 = 40
	var e int64 = 50

	// Unsigned integer types
	var ua uint = 60
	var ub uint8 = 70
	var uc uint16 = 80
	var ud uint32 = 90
	var ue uint64 = 100

	// Floating point types
	var f32 float32 = 1.23
	var f64 float64 = 4.56

	// Complex types
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i

	// Boolean type
	var flag bool = true

	// String type
	var str string = "Hello, Go!"

	// Print all variables
	fmt.Println("Integer types:", a, b, c, d, e)
	fmt.Println("Unsigned integer types:", ua, ub, uc, ud, ue)
	fmt.Println("Floating point types:", f32, f64)
	fmt.Println("Complex types:", c64, c128)
	fmt.Println("Boolean type:", flag)
	fmt.Println("String type:", str)
}