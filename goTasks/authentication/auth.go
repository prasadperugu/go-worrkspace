package main

import (
	"encoding/json"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	// Parse credentials
	credStr := `{"username": "nandu.ahmed@onymos.com", "password": "S0mePa$$w0rd"}`
	var creds Credentials
	err := json.Unmarshal([]byte(credStr), &creds)
	if err != nil {
		fmt.Println("Failed to parse credentials:", err)
		return
	}

	// Check credentials
	if creds.Username != "nandu.ahmed@onymos.com" || creds.Password != "S0mePa$$w0rd" {
		fmt.Println("Invalid credentials")
		return
	}

	// Generate token
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			Issuer: "Portfolio-API",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingKey := []byte("my_secret_key")
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Println("Failed to generate token:", err)
		return
	}

	// Print result
	result := map[string]interface{}{
		"status":   "OK",
		"token":    signedToken,
		"username": creds.Username,
	}
	resultStr, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Failed to marshal result:", err)
		return
	}
	fmt.Println(string(resultStr))
}
