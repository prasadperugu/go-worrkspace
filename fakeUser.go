package main

import (
	"encoding/json"
    "fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//1.Go programming language allows struct fields to be tagged with metadata, 
// which can be used to control the serialization and deserialization of structs when encoding and decoding 
//JSON, XML, or other formats.(encoding/json)
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
type User struct {
	Gender string `json:"gender"`
	Name struct {
		First string `json:"first"`
		Last string `json:"last"`
	} `json:"name"`
	Email string `json:"email"`
}

func main() {
//2.t is a code snippet in Go programming language that sets up a function
//to handle HTTP requests at the default route ("/")
//http.HandleFunc is a function from the net/http package
//that binds a function to a specific URL route (the first argument, "/" in this case).

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get name parameter from URL
		// name := r.URL.Query().Get("name")

		// Consume randomuser.me API to get random user
//3.http.Get is a function from the net/http package that sends an HTTP GET request to a specified URL and returns the server's response. 
//The function returns two values: an *http.Response pointer and an error.		
		resp, err := http.Get("https://randomuser.me/api/")
         fmt.Println(resp);
		if err != nil {
			log.Fatalf("failed to get random user: %v", err)
		}
		defer resp.Body.Close()

		// Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("failed to read response body: %v", err)
		}

		// Unmarshal JSON response into User struct
//4.The purpose of this code is to parse the JSON response from an HTTP API into Go structs.
// The json.Unmarshal function from the encoding/json package can be used to unmarshal the JSON response into the data struct.
// The value of the "results" key in the JSON response will be unmarshalled into the Results field of the data struct.
// The resulting data.Results slice will contain a Go representation of the JSON data.		
		var data struct {
			Results []User `json:"results"`
		}
		if err := json.Unmarshal(body, &data); err != nil {
			log.Fatalf("failed to unmarshal response: %v", err)
		}

		// Return first user
		user := data.Results[0]
		fmt.Println("***fakeUserDetails",user);
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user);
	})
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

	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	
}
//go to chrome and search http://localhost:8080/?name=prasad 