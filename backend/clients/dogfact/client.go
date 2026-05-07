package dogfact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
)

type factApiResponse struct{
	Data []factData `json:"data"`
}

type factData struct{
	ID string `json:"id"`
	Type string `json:"type"`
	Attributes attributesData `json:"attributes"`
}

type attributesData struct{
	Body string `json:"body"`
}

func GetFact(ctx context.Context) (string, error){
	url := "https://dogapi.dog/api/v2/facts?limit=1"


	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil{
		return "", fmt.Errorf("error : %w", err)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	
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