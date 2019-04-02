package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
    //"strconv"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// ParseFolder parses files
// more comments here
func ParseFolder(i_folder string, o_folder string, depth int, list *string) {
    if (depth == 0) {
        return
    }

    if _, err := os.Stat(o_folder); os.IsNotExist(err) {
        fmt.Println(o_folder + " created")
        os.MkdirAll(o_folder, os.ModePerm)
    }

	files, err := ioutil.ReadDir(i_folder)
	if err != nil {
		log.Fatal(err)
	}
   *list += "<ul>" + o_folder[3:]
    for _, f := range files {
            if strings.HasSuffix(f.Name(), ".go") {
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

                out := ""
                out += htmlHeader
                fmt.Println("Parsing:" + f.Name())
                out += "<!-- " + f.Name() + " -->\n"
                out += ParseFile(i_folder + "/" + f.Name())
                out += htmlFooter
                d1 := []byte(out)
                fmt.Println(i_folder+"/"+f.Name()+".html")
                err := ioutil.WriteFile(o_folder+"/"+f.Name()+".html", d1, 0644)
                check(err)
                *list += "<li><a href="+o_folder+"/"+f.Name()+".html>"+f.Name()+"</a></li>"
            } else {
                sub_path := "/" + f.Name()
                i_folder_sub := i_folder + sub_path
                o_folder_sub := o_folder + sub_path
                fi, err := os.Stat(i_folder_sub);
                check(err)
                if fi.Mode().IsDir() {
                    ParseFolder(i_folder_sub, o_folder_sub ,depth-1, list);
                }
            }
    }
   *list += "</ul>"
}
