package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome try to /hello?name=sangam"))
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")

	if name == ""{
		name = "Guest"
	}

	w.Write([]byte(name))
}

func main(){

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)


	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}