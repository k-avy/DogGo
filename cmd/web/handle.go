package main

import (
	"fmt"
	"github.com/kavyashukla26/DogGo/pkg/api"
	"net/http"
	"text/template"
)

type Url string

func (u Url) GetImage() string {
	u = Url(api.Getdata())
	return string(u)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.NotFound(w, r)
		return
	}
	files := []string{"./ui/html/home.page.tmpl", "./ui/html/base.layout.tmpl"}
	templ, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Errorf("error is: %v", err)
		return
	}

	er := templ.Execute(w, Url(api.Getdata()))
	if er != nil {
		http.Error(w, er.Error(), 500)
		fmt.Errorf("error is: %v", err)
		return
	}
}
