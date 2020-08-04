package main

import (
	"flag"
	"fmt"
)

func main() {
	// flag.* arguments are (<flag name>, <default val>, short description)
	basicFlag := flag.String("t", "dev", "Target envrionment")
	flag.Parse()
	fmt.Println("Basic Args `t` flag: " + *basicFlag)
	fmt.Println(flag.Args())
	fmt.Println(flag.)

}
