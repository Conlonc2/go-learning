package main

import (
	"flag"
	"fmt"

	"github.com/bigkevmcd/go-configparser"
)

type config struct {
	port int `yaml:"server.port"`
}

func errMsg() {
	fmt.Println("Error")
	fmt.Println("Usage: main -f <FILE TO READ>")
}

func main() {
	fmt.Println("Application Start...")
	fileName := flag.String("f", "", "File to parse")
	flag.Parse()
	if *fileName == "" {
		errMsg()
		return
	}
	fmt.Printf("File Chosen: %s\n", *fileName)

	cfg, err := configparser.NewConfigParserFromFile("application.properties")
	if err != nil {
		panic("File opporation panic")
	}

	serverPort, _ := cfg.Get("SERVER", "prt")
	testMsg, _ := cfg.Get("TEST", "msg")

	fmt.Println(serverPort)
	fmt.Println(testMsg)

}
