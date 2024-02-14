package main

import (
	"flag"
	"godukto/dukto"
	"log"
)

func main() {
	host := flag.String("h", "192.168.1.195", "set host")
	file := flag.String("f", "", "path to file")

	flag.Parse()

	if err := dukto.SendFile(*file, *host); err != nil {
		log.Fatal(err)
	}
	// dukto.UdpBroadcastListen()
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

// package main
//
// import "fmt"
//
// func main() {
// 	m := make(map[string]string)
//
// 	m["Laptop"] = "192.168.1.195:4644"
// 	m["Laptop"] = "192.168.1.195:4644"
// 	m["Laptop"] = "192.168.1.195:4644"
// 	m["Laptop"] = "192.168.1.195:4644"
// 	m["android"] = "192.168.1.149:4644"
// 	m["android"] = "192.168.1.149:4644"
// 	m["android"] = "192.168.1.149:4644"
// 	m["android"] = "192.168.1.149:4644"
// 	fmt.Println(m)
// }
