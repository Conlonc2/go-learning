package main

import (
	"log"
	"os"
)

func main() {
	logFile := "app.log"
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0655)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer f.Close()
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(f)
	log.Println("Is this log working?")
	log.Println("It sure is...")
	log.Printf("Go look at %s", logFile)
}
