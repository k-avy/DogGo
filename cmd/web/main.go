package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", HomePage)
	err := http.ListenAndServe("127.0.0.1:8888", mux)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
}
