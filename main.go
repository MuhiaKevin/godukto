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

// var wg sync.WaitGroup
// 	// 192.168.1.149
// 	ips := []string{"192.168.1.195", "192.168.1.149"}

// 	for _, item := range ips {
// 		wg.Add(1)

// 		go func(ipAdd string) {
// 			dukto.SendFile(*file, ipAdd)
// 			wg.Done()
// 		}(item)
// 	}

// 	wg.Wait()
