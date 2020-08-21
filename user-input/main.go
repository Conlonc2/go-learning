package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Print("Enter Name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("The name you entered is: %s\n", name)
	log.Printf("The name you entered is: %s\n", name)
}
