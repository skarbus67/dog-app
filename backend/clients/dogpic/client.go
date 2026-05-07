package dogpic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type pictureApiResponse struct{
	Message string `json:"message"`
	Status string `json:"status"`
}

type Client struct{
	httpClient *http.Client
}

func NewClient(c *http.Client) *Client{
	return &Client{httpClient: c}
}

func (c *Client) GetPicture(ctx context.Context) (string, error) {
	url := "https://dog.ceo/api/breeds/image/random"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil{
		return "", fmt.Errorf("error : %w", err)
	}

	client := c.httpClient

	response, err := client.Do(req)
	
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