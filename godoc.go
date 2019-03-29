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
    htmlHeader += `<style> 
    ul{
        list-style:none;
    padding: 5px 0px 5px 5px;
    margin-bottom: 5px;
    border-bottom: 1px solid #efefef;
    font-size: 12px;
    }
           `


    htmlHeader += "</style></head><body>"
    htmlFooter := "</body></html>"
    out += htmlHeader
	ParseFolder(input, output, int(depth), &out)
    out += htmlFooter
    d1 := []byte(out)
    err := ioutil.WriteFile(output+"/index.html", d1, 0644)
    check(err)


    fmt.Println("Documenting: " + input + " into folder: " + output)
}
