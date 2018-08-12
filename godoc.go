package main

import (
	"io/ioutil"
	//	"bufio"
	"fmt"
	//	"io"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// comment
// another comment
func main() {
	input := os.Args[1]
	output := os.Args[2]

	out := ParseFolder(input)
	d1 := []byte(out)
	err := ioutil.WriteFile(output+"/go-doc.index.html", d1, 0644)
	check(err)
	fmt.Println("Documenting: " + input + " into folder: " + output)
}
