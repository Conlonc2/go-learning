// Read a json object (from example.json) and create struct on fly
// Use case: if many formats of json inbound we can loop over each mapping type to create a struct used for dynamic parsing later
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	x := struct {
		Name string
		T    string
	}{}
	file, _ := ioutil.ReadFile("example.json")
	_ = json.Unmarshal([]byte(file), &x)
	fmt.Println(x)
}
