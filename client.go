package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// MakeRequest makes request
func MakeRequest(host string, port int, name string) ([]byte, error) {
	url := fmt.Sprintf("http://%s:%d/hi?name=%s", host, port, name)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil

}

func run() {
	body, err := MakeRequest("localhost", 9999, "james")
	if err != nil {
		fmt.Println("Err--", err)
	}
	fmt.Println(string(body))

}
