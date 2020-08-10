package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/syumai/protocat"
)

var imports = flag.String("I", "", "set import roots using comma separated syntax: -I=.,vendor/googleapis,vendor/google/protobuf")

func main() {
	flag.Parse()
	fileNames := flag.Args()
	importPaths := strings.Split(*imports, ",")
	f, err := protocat.ConcatFiles(importPaths, fileNames...)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		log.Fatal(err)
	}
}
