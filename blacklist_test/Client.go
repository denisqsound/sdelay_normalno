package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/pkg/errors"
	"io/ioutil"

	//"io/ioutil"
	"log"
	//"net/http"
	//"strings"
)

const baseUrl = "http://172.25.84.30:9752"

type AuthResp struct {
	StatusCode int  `json:"status"`
	Body       Body `json:"body"`
}

type Body struct {
	Token string `json:"token"`
}

func main() {

	login_url := fmt.Sprintf("%s/v1/login", baseUrl)

	// BODY
	//	payload := strings.NewReader(`{
	//    "username": "dev@wallarm.com",
	//    "password": "dojsad_knar"
	//}`)

	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Content-Type": "application/json"},
		Data: map[string]string{
			"username": "dev@wallarm.com",
			"password": "dojsad_knar",
		},
	}

	res, err := grequests.Post(login_url, ro)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	body, err := ioutil.ReadAll(res.RawResponse.Body)
	fmt.Printf("BODY : %T", body)

	resp := &AuthResp{}
	if err := json.Unmarshal(body, resp); err != nil {
		log.Fatal(errors.Wrap(err, "json.Unmarshal"))
	}

	authCookies := res.RawResponse.Cookies()[0]

	fmt.Printf("TOKEN : %v", resp.Body.Token)
	fmt.Printf("AUTH COOKIES : %v", authCookies)

	// TODO: Old requests realisations

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//client := &http.Client{}
	//req, err := http.NewRequest(http.MethodPost, url, payload)
	//if err != nil {
	//	//fmt.Println("Fatal")
	//	log.Fatal(errors.Wrap(err, "http.NewRequest"))
	//	return
	//}
	//req.Header.Add("Content-Type", "application/json")
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res.Cookies())
	//authCookies := res.Cookies()[0]
	//defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
	//resp := &AuthResp{}
	//if err := json.Unmarshal(body, resp); err != nil {
	//	log.Fatal(errors.Wrap(err, "json.Unmarshal"))
	//}

	//fmt.Printf("%+v", resp)
	//fmt.Println(resp.Body.Token)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//client := &http.Client{}
	//var addBlacklist string
	//addBlacklist = fmt.Sprintf("%s/v3/blacklist/bulk", baseUrl)
	//
	//req, err := http.NewRequest(http.MethodPost, addBlacklist, payload)
	//if err != nil {
	//	//fmt.Println("Fatal")
	//	log.Fatal(errors.Wrap(err, "http.NewRequest"))
	//	return
	//}
	//req.Header.Add("Content-Type", "application/json")
	//req.AddCookie(authCookies)
	//
	//res, err = client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res.Cookies())
	//
	//defer res.Body.Close()

}
