package main

import (
	"flag"
	"karhuhelsinki/goltrane/convert"
	"karhuhelsinki/goltrane/internal"
	"os"
)

func main() {

	method := os.Args[1:][0]

	infile := flag.String("infile", "index.css", "Specify infile")
	outfile := flag.String("outfile", "outfile.css", "Specify outfile")
	// alpha := flag.String("alpha", "1", "Specify alpha channel")

	flag.Parse()

	data, err := os.ReadFile(*infile)
	internal.CheckErr(err)

	translated := convert.Call(method, string(data))

	err = os.WriteFile(*outfile, []byte(translated), 0644)
	internal.CheckErr(err)
}
