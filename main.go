package main

import (
	"godukto/cmd"
	// "godukto/dukto"
	// "log"
	// "net"
)

func main() {
	cmd.Execute()
}

// func main() {
// 	host := flag.string("h", "192.168.1.195", "set host")
// 	file := flag.string("f", "", "path to file")
//
// 	flag.parse()
//
// 	if err := dukto.sendfile(*file, *host); err != nil {
// 		log.fatal(err)
// 	}
// 	// dukto.udpbroadcastlisten()
// }

// var wg sync.waitgroup
// 	// 192.168.1.149
// 	ips := []string{"192.168.1.195", "192.168.1.149"}

// 	for _, item := range ips {
// 		wg.add(1)

// 		go func(ipadd string) {
// 			dukto.sendfile(*file, ipadd)
// 			wg.done()
// 		}(item)
// 	}

// 	wg.wait()

// package main
//
// import "fmt"
//
// func main() {
// 	m := make(map[string]string)
//
// 	m["laptop"] = "192.168.1.195:4644"
// 	m["laptop"] = "192.168.1.195:4644"
// 	m["laptop"] = "192.168.1.195:4644"
// 	m["laptop"] = "192.168.1.195:4644"
// 	m["android"] = "192.168.1.149:4644"
// 	m["android"] = "192.168.1.149:4644"
// 	m["android"] = "192.168.1.149:4644"
// 	m["android"] = "192.168.1.149:4644"
// 	fmt.println(m)
//
