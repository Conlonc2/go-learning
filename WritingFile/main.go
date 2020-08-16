package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Writing to file
	err := ioutil.WriteFile("test.log", []byte("My first Go file (gfile!)\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Appending to the above file
	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString("New Line Baby!"); err != nil {
		log.Fatal(err)
	}

	// Reading file
	data, err := ioutil.ReadFile("test.log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}
