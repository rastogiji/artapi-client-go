package art

import (
	"net/http"
	"time"
)

const (
	BaseURL = "https://crudcrud.com/api/cfade8eebe3a4b58b9161d89dbd38f65/art"
)

type ArtClient struct {
	BaseURL string
	Client  *http.Client
}

type ArtResp struct {
	ID     string `json:"_id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

type ArtReq struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
}

func NewClient() *ArtClient {
	return &ArtClient{
		BaseURL: BaseURL,
		Client: &http.Client{
			Timeout: time.Minute,
		},
	}
}
