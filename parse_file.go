package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
    "strconv"
)

// todo avoud double highlight.
func highlight(mline string) string {
	var builtins []string
	builtins = append(builtins, "ComplexType")
	builtins = append(builtins, "FloatType")
	builtins = append(builtins, "IntegerType")
	builtins = append(builtins, "Type")
	builtins = append(builtins, "Type1")
	builtins = append(builtins, "bool")
	builtins = append(builtins, "byte")
	builtins = append(builtins, "complex128")
	builtins = append(builtins, "complex64")
	builtins = append(builtins, "error")
	builtins = append(builtins, "float32")
	builtins = append(builtins, "float64")
	builtins = append(builtins, "int")
	//builtins = append(builtins, "int16")
	//builtins = append(builtins, "int32")
	//builtins = append(builtins, "int64")
	//builtins = append(builtins, "int8")
	builtins = append(builtins, "rune")
	builtins = append(builtins, "string")
	builtins = append(builtins, "uint")
	builtins = append(builtins, "uint16")
	builtins = append(builtins, "uint32")
	builtins = append(builtins, "uint64")
	builtins = append(builtins, "uint8")
	builtins = append(builtins, "uintptr")

	for i := range builtins {
		link := " <a href='https://golang.org/pkg/builtin/#" + builtins[i] + "'>" + builtins[i] + "</a>"
		mline = strings.Replace(mline, " "+builtins[i], link, -1)
	}
	return mline
}

// ParseFile parses files
// more comments here
func ParseFile(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	output := ""
	if err != nil {
		fmt.Println("Sorry, I couldn't open the file:" + fileName)
		return ""
	}
	lines := strings.Split(string(dat), "\n")
	previousComments := ""
	fileInfo := ""

	re := regexp.MustCompile("^func(\\s+\\([^)]*\\))* ")
	re2 := regexp.MustCompile("\\(.*")
	funcIndex := ""
	for i := range lines {
		line := lines[i]
		if strings.HasPrefix(line, "//") {
			line = line[2:]
			previousComments += line + "\n<br/>"
		}
		if strings.HasPrefix(line, "func") {
			fname := []byte(line)
			fname = re.ReplaceAll(fname, []byte(""))
			fname = re2.ReplaceAll(fname, []byte(""))

			mline := strings.Replace(line, "{", "", -1)
			mline = highlight(mline)
			foutput := "" 
			foutput += "<a name='" + string(fname) + "'></a><h2>" + string(fname) + "</h2>\n"
			foutput += "<div class=blk>\n"
			foutput += "<code>" + mline + "</code>\n"
			foutput += "<p>\n"
			foutput += previousComments
			foutput += "<a href=https://github.com/NAKsir-melody/go-ethereum/tree/master/"+fileName[14:]+"#L" + strconv.Itoa(i)+">"+fileName[14:]+"#L" + strconv.Itoa(i)+"</a><br>"
			foutput += "</p>\n"
			foutput += "</div>\n"
			previousComments = ""
			// if name is exported, we document it, othwerise, it is subject to change.
		//	if fname[0] >= 'A' && fname[0] <= 'Z' {
				output += foutput
				funcIndex += "<li><a href=\"#" + string(fname) + "\">" + string(fname) + "</a></li>"
				fmt.Println("found method:" + string(fname))
		//	} else {
		//		fmt.Println("found undocumented method:" + string(fname))
		//	}

		}
		if strings.HasPrefix(line, "import") {
			fileInfo = previousComments
			previousComments = ""
		}
	}
	funcIndex = "<ul>" + funcIndex + "</ul>"
	fname := path.Base(fileName)
	header := "<h1>" + fname + "</h1>" + fileInfo + "<br/><br/>" + funcIndex + "<br/><br/>"
	return header + output

}
