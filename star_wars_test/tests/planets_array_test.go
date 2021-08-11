package planets

import (
	. "github.com/denisqsound/sdelay_normalno/star_wars_test/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//var Cases = []interface{}{
//	-1, 0, "qwe", "§", "union+select", "ё",
//}

type Case struct {
	NumberPlanet interface{}
	Expected     Result
}

type Result struct {
	StatusCode int
	Body       string
	Name       string
}

var testCasesPLanet = map[string]Case{
	"case 1": {
		NumberPlanet: -1,
		Expected: Result{
			StatusCode: http.StatusNotFound,
			Body:       "{\"detail\":\"No found\"}",
			Name:       ""},
	},
}

func TestArrayPlanets(t *testing.T) {
	//Arrange
	//PositiveNumbersOfPlanet := []int{1, 3, 7, 15}
	//Expected := []string{"Tatooine", "Yavin IV", "Endor", "Polis Massa"}

	//Act
	for caseName, item := range testCasesPLanet {
		t.Log(caseName)
		planet, status, body := GetPlanet(item.NumberPlanet)

		assert.Equal(t, item.Expected, Result{StatusCode: status, Body: body, Name: planet.Name})

		//if status != item.StatusCode {
		//	t.Errorf("Wrong status code. Expected %d, got %d", item.StatusCode, status)
		//}
		//if planet.Name != item.Name {
		//	t.Errorf("Wrong planet name. Expected %v, got %v", item.Name, planet.Name)
		//}
		//if body != item.Body {
		//	t.Errorf("Wrong body. Expected %v, got %v", item.Body, body)
		//}

	}

}
