package main

import (
	"flag"
	"godukto/dukto"
)

func main() {
	file := flag.String("f", "", "path to file")
	flag.Parse()

	if *file == "" {
		panic("Enter a file")
	}

	dukto.SendFile(*file)
}
