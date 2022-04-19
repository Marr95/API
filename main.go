package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var title string //voer hier je de titel van de film in

func init() {
	flag.StringVar(&title, "Filmtitle", "", "Vul de titel van de film in: ")
	flag.Parse()
}
func main() {

	key := ""
	data, err := ioutil.ReadFile("omd.key") //hier lees jij je key file mee
	if err != nil {
		log.Fatal("Het lezen van deze key is niet mogelijk")
	}
	key = string(data)

	baseUrl := "http://www.omdbapi.com/" //dit is de url van de server
	url := baseUrl + "?t=" + title       //hier voor je titel in op de site
	url += "&apikey=" + key              //hier voor je url plus de api key

	response, err := http.Get(url) // hieronder gaat die kijken of je een repsonse krijgt van de api of niet
	if err != nil {
		log.Fatal("kon de api niet pakken:", err)
	}
	defer response.Body.Close() // hier in word verteld dat die hem straks gaat sluiten en niet nu hij is nog bezig met lezen.

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("kon de response body niet lezen", err)
	}

	var foundmovie Film
	err = json.Unmarshal(data, &foundmovie) // hier pakt hij die json file soort van uit
	if err != nil {
		log.Fatal("kon de json niet lezen:", err)
	}
	fmt.Println("De titel van de film is:", foundmovie.Title)
	fmt.Println("De year van de film is:", foundmovie.Year)
	fmt.Println("De country van de film is:", foundmovie.Countries)
	fmt.Println("De genre van de film is:", foundmovie.Genre)
}

type Film struct {
	Title     string
	Year      string
	Countries string `json:"country"` // dit is referentie naar de json file
	Genre     string
}
