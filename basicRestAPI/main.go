package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var responses []jsonResponse

type jsonResponse struct {
	Title string `json:"Title"`
	Body  string `json:"Msg"`
}

func index(write http.ResponseWriter, req *http.Request) {
	fmt.Fprint(write, "Server is Up!")
	fmt.Println("Index hit")

}
func getResponse(write http.ResponseWriter, req *http.Request) {
	fmt.Println("Retriving response")
	json.NewEncoder(write).Encode(responses)
}

func reqHandler() {
	http.HandleFunc("/", index)
	http.HandleFunc("/getResponse", getResponse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	responses = []jsonResponse{
		jsonResponse{Title: "Title 1", Body: "Body 1"},
		jsonResponse{Title: "Title 2", Body: "Body 2"},
	}
	reqHandler()
}
