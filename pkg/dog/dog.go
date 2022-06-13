package dog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Dogimg struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Getdata() string {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Errorf("error is: %v", err)
		return "Not Found"
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error is: %v", err)
		return err.Error()
	}

	var imgurl Dogimg
	json.Unmarshal(data, &imgurl)
	return imgurl.Message
}

