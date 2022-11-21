package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Peoples struct {
	Count   int            `json:"count"`
	Next    string         `json:"next"`
	Results []SwapiPeoples `json:"results"`
}

type SwapiPeoples struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	Films     []string `json:"films"`
}

func main() {
	response, err := http.Get("https://swapi.dev/api/people")
	if err != nil {
		log.Fatal("error consume api")
	}

	responseBody, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var warga Peoples
	errUnmarshal := json.Unmarshal(responseBody, &warga)
	if errUnmarshal != nil {
		log.Fatal("error unmarshal")
	}

	// fmt.Println("Nama", warga.Results[1].Name)
	// fmt.Println("hair color:", orang1.HairColor)
	// fmt.Println("skin color:", orang1.SkinColor)

	for _, v := range warga.Results {
		fmt.Println("Nama:", v.Name)
		fmt.Println("HairColor:", v.HairColor)
	}
}
