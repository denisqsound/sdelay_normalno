package planets

import (
	"fmt"
	. "github.com/denisqsound/sdelay_normalno/star_wars_test/models"
	"github.com/levigross/grequests"
	"log"
)

const BaseUrl = "https://swapi.dev/api/"

func GetPlanet(PlanetNumber int) *Planet {
	url := fmt.Sprintf("%s/planets/%d", BaseUrl, PlanetNumber)
	res, err := grequests.Get(url, nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	resp := Planet{}
	err = res.JSON(&resp)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return &resp

}
