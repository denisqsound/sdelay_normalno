package tools

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/sirupsen/logrus"
	"gitlab.ozon.ru/auth/login-tool/pkg/login"
	"golang.org/x/oauth2"
	"log"
)

func main() {
	data := ""
	sDec, _ := b64.StdEncoding.DecodeString(data)
	rty := string(sDec)

	data2 := ""
	sDec2, _ := b64.StdEncoding.DecodeString(data2)
	fgh := string(sDec2)

	var ts oauth2.TokenSource
	var err error
	err = retry.Do(
		func() error {
			ts, err = login.Login(login.Options{
				AuthConfig:   "auth-config.auth.stg.s.o3.ru",
				ClientSecret: "",
				ClientId:     "internal-projects.staff-api",
				Username:     fgh,
				Password:     rty,
				Scopes:       []string{"profile roles"},
			})
			if err != nil {
				logrus.Errorf("Can't get bearer token:", err)
				return err
			}
			return nil
		})

	token, err := ts.Token()
	if err != nil {
		log.Fatalln("Can't get execute Token method : ", err)
	}
	if token == nil {
		log.Fatalln("Token should be not nil : ", token)
	}

	tokenHeader := fmt.Sprintf("Bearer %s", token.AccessToken)

	fmt.Printf("%v", tokenHeader)

}
