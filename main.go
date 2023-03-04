package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/thanhpk/randstr"
	"fmt"
	"math/rand"
	"os"
)

var sticky string

func RootEndpoint(w http.ResponseWriter, req *http.Request){
	fmt.Println(sticky)
	json.NewEncoder(w).Encode(sticky)

}

func TokenEndpoint(w http.ResponseWriter, req *http.Request){
	token := os.Getenv("TOKEN")
	fmt.Println(token)
	json.NewEncoder(w).Encode(token)
}

func ErrorEndpoint(w http.ResponseWriter, req *http.Request){
	num := rand.Intn(5)
	fmt.Println(num)
	if num == 0 {
		os.Exit(1)
	}
	json.NewEncoder(w).Encode(num)
}

func main(){
	sticky = randstr.String(3)

	router := mux.NewRouter()
	// routes
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	router.HandleFunc("/token", TokenEndpoint).Methods("GET")
	router.HandleFunc("/error", ErrorEndpoint).Methods("GET")

	// start server
	log.Fatal( http.ListenAndServe(":3001", router) )

}