package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok" : true,
		"message" : "Json encode successfully",
		"datetime" : time.Now().UTC(),
	} 

	err := json.NewEncoder(w).Encode(res)
	fmt.Println(err)
}

func main(){
	http.HandleFunc("/ok", successHandler)
	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}