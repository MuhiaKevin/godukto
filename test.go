package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["Laptop"] = "192.168.1.195:4644"
	m["Laptop"] = "192.168.1.195:4644"
	m["Laptop"] = "192.168.1.195:4644"
	m["Laptop"] = "192.168.1.195:4644"
	m["android"] = "192.168.1.149:4644"
	m["android"] = "192.168.1.149:4644"
	m["android"] = "192.168.1.149:4644"
	m["android"] = "192.168.1.149:4644"
	fmt.Println(m)
}
