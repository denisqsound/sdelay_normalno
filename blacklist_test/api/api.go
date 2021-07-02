package api

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/pkg/errors"
	"log"
	"os"
)

const (
	BaseUrl         = "api://172.25.84.30:9752"
	LoginUrl        = "%s/v1/login"
	AddBlacklistUrl = "%s/v3/blacklist/bulk"
)

var jsonContentType = map[string]string{"Content-Type": "application/json"}

type BlacklistApi struct {
	user *AuthorisedUser
}

func (a *BlacklistApi) getAuth() {
	authUrl := fmt.Sprintf(LoginUrl, BaseUrl)
	Username, _ := os.LookupEnv("USERNAME")
	Password, _ := os.LookupEnv("PASSWORD")

	ro := &grequests.RequestOptions{
		Headers: jsonContentType,
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
	a.user = &AuthorisedUser{
		Token:  resp.Body.Token,
		Cookie: authCookies,
	}

}

func (a *BlacklistApi) AddToBlacklist(bl []Blacklist) error {
	if a.user != nil {
		a.getAuth()
	}
	url := fmt.Sprintf(AddBlacklistUrl, BaseUrl)
	ro := &grequests.RequestOptions{
		Headers: jsonContentType,
		JSON:    AddBlacklistReq{Bulk: bl},
	}

	res, err := grequests.Post(url, ro)
	//TODO Поменять на errors.Wrap
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	resp := AddBlacklistResp{}
	err = res.JSON(&resp)
	if err != nil {
		return errors.Wrap(err, "Unable to make request: ")
	}
	return nil
}
