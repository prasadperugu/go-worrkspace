package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Name struct {
	Name string `json:"name"`
}

type Response struct {
	Gender string `json:"gender"`
	Name struct {
		First string `json:"first"`
		Last string `json:"last"`
	} `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/myapp", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var name Name
			err := json.NewDecoder(r.Body).Decode(&name)
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			response := Response{
				Gender: "male",
				Name: struct {
					First string `json:"first"`
					Last string `json:"last"`
				}{
					First: "Erol",
					Last:  "Diehl",
				},
				Email: "erol.diehl@example.com",
			}

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("prasad");

	log.Fatal(http.ListenAndServe(":8080", nil))
//The main function sets up an HTTP handler for the /myapp endpoint that listens for POST requests and returns a JSON response.
}
