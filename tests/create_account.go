package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/v1/user/create"
	method := "POST"

	s := `{
    "user_name":"name_%d",
    "password":"password",
    "user_type":1,
    "email":"name_%d@gmail.com",
    "profile_uri":"profile_uri"
}`
	var p string
	for i := 0; i < 10000; i++ {
		p = fmt.Sprintf(s, i, i)
		fmt.Println(p)
		payload := strings.NewReader(p)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	}
}
