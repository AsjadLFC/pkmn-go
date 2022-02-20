package main

import (
	"encoding/json"
	"fmt"
	"go-test-http/handlers"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var baseData BaseData
var descriptionBD []Description
var pokemonBD []Pokemon
var versionBD []Version

type BaseData struct {
	Descriptions []Description `json:"descriptions"`
	Name         string        `json:"name"`
	Pokemons     []Pokemon     `json:"pokemon_entries"`
	Versions     []Version     `json:"version_groups"`
}

type Pokemon struct {
	Entry   int `json:"entry_number"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Description struct {
	Description string `json:"description"`
	Language    string `json:"language"`
}

func main() {
	jsonFile, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(jsonFile, &baseData)

	descriptionBD = baseData.Descriptions
	pokemonBD = baseData.Pokemons
	versionBD = baseData.Versions

	r := gin.Default()

	// List all Pokemon
	r.GET("/get", handlers.GetMonster())

	// Search Pokemon
	// r.POST("/search", handlers.SearchPKMN())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
