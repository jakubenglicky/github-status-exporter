package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Response struct {
    Components []Component `json:"components"`
}

type Component struct {
	Name string `json:"name"`
	Status string `json:"status"`
}


func GetGithubStatusComponents() []Component {

	resp, err := http.Get("https://www.githubstatus.com/api/v2/summary.json")

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var result Response
    err = json.Unmarshal(data, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result.Components
}
