package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type factApiResponse struct{
	Data []FactData `json:"data"`
}

type FactData struct{
	ID string `json:"id"`
	Type string `json:"type"`
	Attributes AttributesData `json:"attributes"`
}

type AttributesData struct{
	Body string `json:"body"`
}

func GetFact() (string, error){
	url := "https://dogapi.dog/api/v2/facts?limit=1"

	response, err := http.Get(url)

	if err != nil{
		return "", fmt.Errorf("failed fetching the dog fact : %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200{
		return "", fmt.Errorf("wrong status code : %v", response.StatusCode)
	}

	var result factApiResponse

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return "", fmt.Errorf("failed decoding json : %w", err)
	}

	return result.Data[0].Attributes.Body, nil

}