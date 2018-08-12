package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// ParseFolder parses files
// more comments here
func ParseFolder(folderName string) string {
	out := ""
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}
	htmlHeader := "<!DOCTYPE html>\n<html><head><title> go-doc Documentation</title>"
	htmlHeader += "<style> "
	htmlHeader += "* { font-family:helvetica; }\n"
	htmlHeader += ".blk {padding:10px; }\n"
	htmlHeader += "h1,h2,h3 {color:darkblue;}  \n"
	htmlHeader += "p {font-family:\"Courier New\",monospace;}  \n"
	htmlHeader += "code { padding:10px; background-color:rgb(230,230,230); display:block; margin-bottom:10px; }\n"
	//	htmlHeader += ". blk {padding:10px};\n"
	//	htmlHeader += ". blk {padding:10px};\n"
	htmlHeader += "</style></head>"
	htmlFooter := "</body></html>"

	out += htmlHeader
	for _, f := range files {
		fmt.Println(f.Name())

		if strings.HasSuffix(f.Name(), ".go") {
			fmt.Println("Parsing:" + f.Name())
			out += "<!-- " + f.Name() + " -->\n"
			out += ParseFile(folderName + "/" + f.Name())
		}
	}

	out += htmlFooter
	return out

}
