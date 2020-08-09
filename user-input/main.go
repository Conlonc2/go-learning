package main

import "fmt"

func main() {
	fmt.Print("Enter Name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("The name you entered is: %s\n", name)
}
