package planets

import (
	"fmt"
	. "github.com/denisqsound/sdelay_normalno/star_wars_test/api"
	"testing"
)

//type PLanetCases struct {
//
//}

func TestPlanets(t *testing.T) {
	//Arrange
	PositiveNumbersOfPlanet := []int{1, 3, 7, 15}
	Expected := []string{"Tatooine", "Yavin IV", "Endor", "Polis Massa"}
	//NegativeNumbersOfPlanetInt := []int{-1, 0}
	//NegativeNumbersOfPlanetStr := []string{"qwe", "§", "union+select", "ё"}

	//Act
	for _, item := range PositiveNumbersOfPlanet {
		result := GetPlanet(item)
		for _, expectedItem := range Expected {
			if result.Name != expectedItem {
				t.Errorf("Wrong planet name. Expected %v, got %v", Expected, result)
			}
		}
		fmt.Println(result.Name)
	}

}
