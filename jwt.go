package main

import (
	"encoding/json"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Name  string
	Email string
	Role  string
}

var users = []User{
	{
		Name:  "Prasad",
		Email: "prasad.perugu@onymos.com",
		Role:  "admin",
	},
	{
		Name:  "Nandu",
		Email: "nandu.ahmed@onymos.com",
		Role:  "clinician",
	},
}
func getRequestPassword() string {
	// Replace this function with your code to retrieve the password from the request
	return "SomePassword"
}

func getToken(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	password := "SomePassword"
	encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))
	if requestPassword := getRequestPassword(); requestPassword != password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("Authorized")
	fmt.Println("Encoded password:", encodedPassword)

	// Verify the password
	if  request.Password != "encodedPassword" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Search for the user
	var user *User
	for _, u := range users {
		if u.Email == request.Username {
			user = &u
			break
		}
	}
	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "my-api",
		"iat":  time.Now().Unix(),
		"name": user.Name,
		"email": user.Email,
		"role": user.Role,
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func main() {
	http.HandleFunc("/getToken", getToken)
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
