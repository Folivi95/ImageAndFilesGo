package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var mySigningKey = []byte("dummysignedkey")

func main() {
	handleRequests()
}

func GenerateJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "@binator_1308"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("generating JWT Token failed")
		return "", err
	}

	return tokenString, nil
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWTToken()

	// make http request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000", nil)
	req.Header.Set("Token", validToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))
}

func handleRequests() {
	http.HandleFunc("/", HomePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}
