package main

import (
	"flag"
	"godukto/dukto"
)

func main() {
	host := flag.String("h", "192.168.1.195", "set host")
	file := flag.String("f", "", "path to file")

	flag.Parse()

	if *file == "" {
		panic("Enter a file")
	}
	dukto.SendFile(*file, *host)
}
