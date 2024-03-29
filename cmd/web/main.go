package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", HomePage)
	mux.Handle("/", http.FileServer(http.Dir("ui/css/")))
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
}
