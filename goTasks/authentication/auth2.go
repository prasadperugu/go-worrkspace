package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/auth", handleAuth)
	http.ListenAndServe(":8080", nil)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate the user (replace with your own authentication logic)
	if data.Username != "nandu.ahmed@onymos.com" || data.Password != "S0mePa$$w0rd" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   "Portal-API",
		"iat":   time.Now().Unix(),
		"name":  "Nandu",
		"email": data.Username,
		"role":  "client",
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the token and user information as a JSON response
	response := struct {
		Status   string `json:"status"`
		Token    string `json:"token"`
		Username string `json:"username"`
	}{
		Status:   "OK",
		Token:    tokenString,
		Username: data.Username,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//curl -X POST -H "Content-Type: application/json" -d '{"username": "nandu.ahmed@onymos.com", "password": "S0mePa$$w0rd"}' http://localhost:8080/auth
