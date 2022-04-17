package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Payload struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type createShopResp struct {
	ShopID string `json:"shop_id" `
}

func main() {

	for i := 1000; i < 100000; i++ {

		url := fmt.Sprintf("http://localhost:8080/v1/user/login?user_name=name_%d&password=password&user_type=1", i)
		method := "GET"

		payload := strings.NewReader(``)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		var sessionID string
		for _, v := range res.Cookies() {
			if v.Name == "session_id" {
				sessionID = v.Value
			}
		}
		fmt.Println()
		fmt.Println(sessionID)
		if len(sessionID) == 0 {
			fmt.Println("session empty")
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))

		url = "http://localhost:8080/v1/seller/create_shop"
		method = "POST"

		for j := 1; j < 10000; j++ {
			s := `{
    "name":"name_%d_shop_%d",
    "introduction":"introduction"
}`
			cs := fmt.Sprintf(s, i, j)
			payload = strings.NewReader(cs)

			client = &http.Client{}
			req, err = http.NewRequest(method, url, payload)

			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Cookie", "session_id="+sessionID)

			res, err = client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err = ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			var dateshopResp createShopResp
			result := Payload{Data: &dateshopResp}
			err = json.Unmarshal(body, &result)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(result)
			fmt.Println(string(body))
			shop_id := dateshopResp.ShopID
			fmt.Println(shop_id)
		}
		/*
		   			for z := 1; z < 500; z++ {
		   				url = "http://localhost:8080/v1/seller/create_product"
		   				method = "POST"

		   				s := `{
		       "shop_id": "%s",
		       "title":"title",
		       "cover_uri":"cover_uri",
		       "Price":10,
		       "stock":10,
		       "brand_id":1,
		       "category_id":1,
		       "details":["detail_1", "detail_2", "detail_3"],
		       "show_uris":["uri_1", "uri_2", "uri_3"],
		       "show_uris":["uri_1", "uri_2", "uri_3"],
		       "attr_info":[{"name": "attr_1", "value":"a1"}, {"name": "attr_2", "value":"2"}]
		   }`
		   				shopInfo := fmt.Sprintf(s, shop_id)
		   				client = &http.Client{}
		   				payload = strings.NewReader(shopInfo)
		   				req, err = http.NewRequest(method, url, payload)

		   				if err != nil {
		   					fmt.Println(err)
		   					return
		   				}
		   				req.Header.Add("Content-Type", "application/json")
		   				req.Header.Add("Cookie", "session_id="+sessionID)

		   				res, err = client.Do(req)
		   				if err != nil {
		   					fmt.Println(err)
		   					continue
		   				}
		   				defer res.Body.Close()

		   				body, err = ioutil.ReadAll(res.Body)
		   				if err != nil {
		   					fmt.Println(err)
		   					return
		   				}
		   				fmt.Println(string(body))

		   			}
		   		}

		*/

	}
}
