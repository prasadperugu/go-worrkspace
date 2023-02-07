package main

import (
	"fmt"
	"net/http"
)
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Hospital string `json:"hospital"`
	Manager  string `json:"manager"`
	Department string `json:"department"`
}

func infoAPI(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:     "Nandu",
		Email:    "nandu.ahmed@onymos.com",
		Role:     "clinician",
		Password: "Ear@",
		Hospital: "WWW",
		Manager:  "YYY",
		Department: "medical",
	}


	fmt.Fprintf(w, "Name: %s, Email: %s, Role: %s, Password: %s, Hospital: %s, Manager: %s, Department: %s",
		user.Name, user.Email, user.Role, user.Password, user.Hospital, user.Manager, user.Department)
}

func main() {
	http.HandleFunc("/infoapi", infoAPI)

	http.ListenAndServe(":8080", nil)

}
