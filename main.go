package main

import (
	"log"
	"net/http"
	"jwt-go/jwt"
)

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
