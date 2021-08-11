package planets

import (
	"fmt"
	. "github.com/denisqsound/sdelay_normalno/star_wars_test/api"
	"testing"
)

func TestSimplePlanets(t *testing.T) {
	//Arrange
	PositiveNumbersOfPlanet := 1
	Expected := "Tatooine"

	//Act
	result, _, _ := GetPlanet(PositiveNumbersOfPlanet)
	fmt.Println(result)

	//Assert
	if result.Name != Expected {
		t.Errorf("Wrong planet name. Expected %v, got %v", Expected, result)
	}
}
