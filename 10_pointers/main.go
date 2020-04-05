// a pointer allow to point to a memory address or location of a variable

package main

import "fmt"

func main() {

	a := 5
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read value from address
	fmt.Println(*b)

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}
