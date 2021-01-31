package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var MySigningKey = []byte("mysecret")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf(("There was an error"))
				}
				return MySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
			fmt.Println(token)
		} else {
			fmt.Fprintf(w, "No Authorization Token provided")
		}
	})
}
func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}
func main() {
	fmt.Println("server")
	handleRequests()
}
