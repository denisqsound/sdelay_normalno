package main

import (
	"fmt"
	"log"

	"github.com/levigross/grequests"
)

const (
	BaseUrl = "https://swapi.dev/api/"
	Planets = "%s/planets/"
)

// TODO Как передать номер планеты?
//var jsonContentType = map[string]string{"Content-Type": "application/json"}

func GetPlanet() {
	//url := fmt.Printf(Planets, BaseUrl)
	url := fmt.Sprintf(Planets, BaseUrl)
	res, err := grequests.Get(url)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	resp := Planet{}
	err = res.JSON(&resp)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp

}
