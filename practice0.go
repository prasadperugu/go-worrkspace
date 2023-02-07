package main

import(

 "net/http"
 "fmt"
 "github.com/dgrijalva/jwt-go"
//  go mod tidy

var SECRET = []byte("super_secret-Key");

)
 func CreateJWT() (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.hour).Unix()
	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
        fmt.Println(err.Error())
		return " ",err

	}
	return tokenStr,nil

 }

 func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.handler{
	return http.HandleFunc(func (w http.ResponseWriter, r *http.Request){
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.token)(interface{}, error)){
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				 return SECRET,nil

			}
			if err! = nil {

			}
		}
	}

 }





func Home(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w,"super secret area");
	// fmt.Println(w,"super secret area");


}

func main(){

	http.HandleFunc("/boom",Home)
	
	http.ListenAndServe(":3500",nil)



}