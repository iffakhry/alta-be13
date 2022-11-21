package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SwapiPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	Films     []string `json:"films"`
}

func main() {
	response, err := http.Get("https://swapi.dev/api/people/1")
	if err != nil {
		log.Fatal("error consume api")
	}

	responseBody, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var orang1 SwapiPeople
	errUnmarshal := json.Unmarshal(responseBody, &orang1)
	if errUnmarshal != nil {
		log.Fatal("error unmarshal")
	}

	fmt.Println("Nama", orang1.Name)
	fmt.Println("hair color:", orang1.HairColor)
	fmt.Println("skin color:", orang1.SkinColor)

	for _, v := range orang1.Films {
		fmt.Println("film:", v)
	}
}
