package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Welcome to JWT Application")
	jwtToken, err := GenerateJWTToken()
	if err != nil {
		log.Fatal("Failed to generate token: ", err)
	}

	fmt.Printf("JWT Token: %v", jwtToken)

	handleRequests()
}

func GenerateJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "@binator_1308"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SigningString()

	if err != nil {
		fmt.Errorf("generating JWT Token failed")
		return "", err
	}

	return tokenString, nil
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWTToken()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

func handleRequests() {
	http.HandleFunc("/", HomePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}
