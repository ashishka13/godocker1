package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", welcome).Methods("GET", "PUT", "POST", "DELETE")
	r.HandleFunc("/customers", getmethod).Methods("GET")
	r.HandleFunc("/customers", postmethod).Methods("POST")
	r.HandleFunc("/customer/{custID}", putmethod).Methods("PUT")
	r.HandleFunc("/customer/{custID}", deletemethod).Methods("DELETE")

	fmt.Println("listening..")
	http.ListenAndServe(":5005", r)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to docker rest app, /customers, /customer/{custID}"))
}

func getmethod(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to getmethod"))
}

func putmethod(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to putmethod"))
}

func postmethod(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to postmethod"))
}

func deletemethod(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to deletemethod"))
}
