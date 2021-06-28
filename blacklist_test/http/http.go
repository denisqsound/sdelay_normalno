package http

import (
	"fmt"
	"github.com/levigross/grequests"
	"log"
	"os"
)

func GetRequest() AuthoriseUser {
	authUrl := fmt.Sprintf(LoginUrl, BaseUrl)
	Username, _ := os.LookupEnv("USERNAME")
	Password, _ := os.LookupEnv("PASSWORD")

	ro := &grequests.RequestOptions{
		Headers: map[string]string{"Content-Type": "application/json"},
		JSON: AuthReg{
			Username: Username,
			Password: Password,
		},
	}

	res, err := grequests.Post(authUrl, ro)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	resp := AuthResp{}
	err = res.JSON(&resp)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	authCookies := res.RawResponse.Cookies()[0].Value

	fmt.Printf("BODY : %+v \n", resp.Body.Token)
	fmt.Printf("AUTH COOKIES : %+v \n", authCookies)

	// Записать в структуру токен и куки
	AuthoriseUser := &AuthorisedUser{
		token:  resp.Body.Token,
		cookie: authCookies,
	}

	return AuthoriseUser

}
