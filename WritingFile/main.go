package main

/*
From os page at golang.org

    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
    O_RDONLY int = syscall.O_RDONLY // open the file read-only.
    O_WRONLY int = syscall.O_WRONLY // open the file write-only.
    O_RDWR   int = syscall.O_RDWR   // open the file read-write.
    // The remaining values may be or'ed in to control behavior.
    O_APPEND int = syscall.O_APPEND // append data to the file when writing.
    O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
    O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
    O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
    O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
*/

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
