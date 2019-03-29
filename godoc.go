package main

import (
	"fmt"
	"os"
    "strconv"
	"io/ioutil"
)

// comment
// another comment
func main() {
	input := os.Args[1]
	output := os.Args[2]
	depth, _ := strconv.Atoi(os.Args[3])

    out := ""

    htmlHeader := "<!DOCTYPE html>\n<html><head><title> go-doc Documentation</title>"
    htmlHeader += "<style> "
    htmlHeader += "* { font-family:helvetica; }\n"
    htmlHeader += ".blk {padding:10px; }\n"
    htmlHeader += "h1,h2,h3 {color:darkblue;}  \n"
    htmlHeader += "p {font-family:\"Courier New\",monospace;}  \n"
    htmlHeader += "code { padding:10px; background-color:rgb(230,230,230); display:block; margin-bottom:10px; }\n"
    //	htmlHeader += ". blk {padding:10px};\n"
    //	htmlHeader += ". blk {padding:10px};\n"
    htmlHeader += "<meta charset='UTF-8'>"
    htmlHeader += "</style></head>"
    htmlFooter := "</body></html>"

    out += htmlHeader

	ParseFolder(input, output, int(depth), &out)

    out += htmlFooter
    d1 := []byte(out)
    err := ioutil.WriteFile(output+"/index.html", d1, 0644)
    check(err)


    fmt.Println("Documenting: " + input + " into folder: " + output)
}
