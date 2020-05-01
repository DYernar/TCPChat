package main

import(
	"os"
	"fmt"
)


func main() {
	args := os.Args

	if len(args) == 2 {
		CreateServer(args[1])
	} else if len(args) == 1 {
		CreateServer("")
	} else {
		fmt.Println("[USAGE]: ./TCPChat $port")
	}
}