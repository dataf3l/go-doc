package main

import (
	"fmt"
	"os"
    "strconv"
)

// comment
// another comment
func main() {
	input := os.Args[1]
	output := os.Args[2]
	depth, _ := strconv.Atoi(os.Args[3])

	ParseFolder(input, output, int(depth))
    fmt.Println("Documenting: " + input + " into folder: " + output)
}
