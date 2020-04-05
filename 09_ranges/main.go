// Range is used to loops throw arrays, maps ...

package main

import "fmt"

func main() {
	var ids = []int{32, 45, 33, 65, 75}

	// Using index
	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }

	// Not using index

	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	// Add ids together
	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Printf("Sum: %d\n", sum)

	// Range with map

	var emails = map[string]string{
		"bob":   "bob@gmail.com",
		"alice": "alice@gmail.com",
		"selom": "selom@gmail.com",
	}

	for name, email := range emails {
		fmt.Printf("%s: %s\n", name, email)
	}
}
