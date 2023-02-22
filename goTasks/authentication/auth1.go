package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User struct represents the user's credentials
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT struct represents the JWT token
type JWT struct {
	Token string `json:"token"`
}

// AuthResponse struct represents the authentication response
type AuthResponse struct {
	Status   string `json:"status"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

func main() {
	http.HandleFunc("/auth", authHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verify the user's credentials
	if user.Username != "nandu.ahmed@onymos.com" || user.Password != "S0mePa$$w0rd" {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := generateToken(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Construct authentication response
	response := AuthResponse{
		Status:   "OK",
		Token:    token.Token,
		Username: user.Username,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func generateToken(username string) (*JWT, error) {
	// Generate JWT token using username
	// This is just an example and should not be used in production
	token := JWT{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJQb3J0YWwtQVBJIiwiaWF0IjoxNTE2MjM5MDIyLCJuYW1lIjoiTmFuZHUiLCJlbWFpbCI6Im5hbmR1LmFobWVkQG9ueW1vcy5jb20iLCJyb2xlIjoiY2xpbmljaWFuIn0.H5JSmevMLjtgX2tEIs3jHrb7bSMSCU9gFyd6vfGm7SU",
	}
	fmt.Println(token)
	return &token, nil
}

//to see the geenerated output:
// curl -X POST -H "Content-Type: application/json" -d '{"username": "nandu.ahmed@onymos.com", "password": "S0mePa$$w0rd"}' http://localhost:8080/auth
//It defines a single HTTP handler function authHandler that handles POST requests to the /auth endpoint.
