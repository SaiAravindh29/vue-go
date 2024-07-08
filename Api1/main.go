package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("Server started .....")
	http.HandleFunc("/validate", validate)
	http.ListenAndServe(":8095", nil)

}

type Validation struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Role string `json:"role"`
}

// var valid []Validation

func validate(w http.ResponseWriter, r *http.Request) {
	log.Println("Validating...")
	// (w).Header().Set("Access-Control-Allow-Origin", "*")
	// (w).Header().Set("Access-Control-Allow-Credentials", "true")
	// (w).Header().Set("Access-Control-Allow-Methods", "GET")
	// (w).Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(valid)
	fmt.Fprint(w, "Hello")
}
