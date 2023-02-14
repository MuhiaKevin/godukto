package main

import "github.com/MuhiaKevin/godukto/networking"

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(2)
// 	go networking.SendBroadcast(&wg)

// 	go networking.SendFile(&wg)

// 	wg.Wait()
// }

func main() {
	networking.SendFile()

}
