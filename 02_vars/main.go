package main

import "fmt"

// const name = "kokou" // globally

func main() {
	// MAIN TYPES
	// string
	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// rune - alias for int32
	// float32 float64
	// complex64 complex128
	// byte - alias for uint8
	// rune - alias for int32

	// Using var
	// var name = "Selom"

	// name := "Selom"
	// size := 1.3

	name, email := "Selom", "shellom2005@gmail.com"
	var age int64 = 37
	const isCool = true

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", email)
}
