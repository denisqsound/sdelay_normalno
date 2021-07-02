package main

import (
	"fmt"
	"testing"
)

func TestPlanets(t *testing.T) {
	//Arrange
	PositiveNumbersOfPlanet := []int{1, 3, 7, 100}
	NegativeNumbersOfPlanetInt := []int{-1, 0}
	NegativeNumbersOfPlanetStr := []string{"qwe", "§", "union+select", "ё"}

	//Act
	//result := GetPlanet()
	fmt.Println(PositiveNumbersOfPlanet, NegativeNumbersOfPlanetInt, NegativeNumbersOfPlanetStr)

}
