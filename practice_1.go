package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		fmt.Printf("%p\n", &slice)
		slice = append(slice, "a")
		fmt.Printf("%p", slice)
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
