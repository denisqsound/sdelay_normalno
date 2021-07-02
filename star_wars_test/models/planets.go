package main

type Planet struct {
	Climate        string   `json:"climate"`
	Created        string   `json:"created"`
	Diameter       string   `json:"diametr"`
	Edited         string   `json:"edited"`
	Films          []string `json:"films"`
	Gravity        string   `json:"gravity"`
	Name           string   `json:"name"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	RotationPeriod string   `json:"rotation_period"`
	SurfaceWater   string   `json:"surface_water"`
	Terrain        string   `json:"terrain"`
	Url            string   `json:"url"`
}

/*
{
    "climate": "Arid",
    "created": "2014-12-09T13:50:49.641000Z",
    "diameter": "10465",
    "edited": "2014-12-15T13:48:16.167217Z",
    "films": [
        "https://swapi.dev/api/films/1/",
        ...
    ],
    "gravity": "1",
    "name": "Tatooine",
    "orbital_period": "304",
    "population": "120000",
    "residents": [
        "https://swapi.dev/api/people/1/",
        ...
    ],
    "rotation_period": "23",
    "surface_water": "1",
    "terrain": "Dessert",
    "url": "https://swapi.dev/api/planets/1/"
}



*/
