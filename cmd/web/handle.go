package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/kavyashukla26/DogGo/pkg/dog"
	"github.com/kavyashukla26/DogGo/pkg/quote"
)

type Url string
type Quo string 


func (u Url) GetImage() string {
	u = Url(dog.Getdata())
	return string(u)
}


func (q Quo) GetQuote() string {
	q = Quo(quote.Getdata())
	return string(q)
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

	er := templ.Execute(w, Url(dog.Getdata()))
	if er != nil {
		http.Error(w, er.Error(), 500)
		fmt.Errorf("error is: %v", err)
		return
	}
}

func QuotePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/quotes" {
		http.NotFound(w, r)
		return
	}
	files := []string{"./ui/html/quotes.page.tmpl", "./ui/html/base.layout.tmpl"}
	templ, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Errorf("error is: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	
	er := templ.Execute(w, Quo(quote.Getdata()))
	if er != nil {
		fmt.Errorf("error is: %v", err)
		http.Error(w, er.Error(), 500)
		return
	}
}
