package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Arti string `json:"arti"`
	// Asma       string `json:"asma"`
	Keterangan string `json:"keterangan"`
}

func main() {
	response, err := http.Get("https://al-quran-8d642.firebaseio.com/data.json?print=pretty")
	if err != nil {
		log.Fatal("error consume api")
	}

	responseBody, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var listData []Data
	errUnmarshal := json.Unmarshal(responseBody, &listData)
	if errUnmarshal != nil {
		log.Fatal("error unmarshal")
	}

	fmt.Println("Arti:", listData[0].Arti)
	// fmt.Println("Asma:", listData[0].Asma)
	fmt.Println("keterangan:", listData[0].Keterangan)
	fmt.Println("-------")

	// for _, v := range listData {
	// 	fmt.Println("Arti:", v.Arti)
	// 	fmt.Println("Asma:", v.Asma)
	// 	fmt.Println("keterangan:", v.Keterangan)
	// 	fmt.Println("-------")
	// }
}
