package dogpic

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type pictureApiResponse struct{
	Message string `json:"message"`
	Status string `json:"status"`
}

func GetPicture() (string, error) {
	url := "https://dog.ceo/api/breeds/image/random"

	response, err := http.Get(url)
	
	if err != nil {
		return "", fmt.Errorf("failed fetching dog picture : %w", err)
	}

	defer response.Body.Close()

	if(response.StatusCode != 200){
		return "", fmt.Errorf("wrong status code : %v", response.StatusCode)
	}

	var result pictureApiResponse

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return "", fmt.Errorf("failed decoding json : %w", err)
	}

	return result.Message, nil

}