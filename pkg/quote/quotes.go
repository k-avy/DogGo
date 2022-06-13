package quote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	Content      string   `json:"content"`
	Id           string   `json:"_id"`
	Tag          []string `json:"tag"`
	Author       string   `json:"author"`
	AuthorSlug   string   `json:"authorslug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateadded"`
	DateModified string   `json:"datemodified"`
}

func Getdata() string {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		fmt.Errorf("error is: %v", err)
		return "Not Found"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error is: %v", err)
		return "nil"
	}

	var quoteurl Quote
	json.Unmarshal(data, &quoteurl)
	return quoteurl.Content

}
