package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Name string `json:"name"`
	price  float    `json: "age"`
}

type People []Person

func allPeople (w http.ResponseWriter, r *http.Request){
	people := People{
		Person{Name : "longld", Age: 22},
	}
	fmt.Println("Endpoint people")
	json.NewEncoder(w).Encode(people)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Long test writer")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/people", allPeople)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}
