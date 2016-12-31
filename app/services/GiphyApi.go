package services

import (
	"GiphyApi/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const GIPHY_API_BETA_KEY = "dc6zaTOxFJmzC"

func GiphyClient() *giphyClient {
	return &giphyClient{
		Mock: false,
	}
}

type giphyClient struct {
	Mock bool
}

func (c *giphyClient) Search(keywords string, respStruct *models.ApiGifs) error {
	if c.Mock {
		return mockSearch(keywords, respStruct)
	} else {
		return search(keywords, respStruct)
	}
}

func search(keywords string, respStruct *models.ApiGifs) error {
	url := fmt.Sprintf(
		"http://api.giphy.com/v1/gifs/search?q=%s&api_key=%s",
		keywords,
		GIPHY_API_BETA_KEY,
	)

	client := &http.Client{Timeout: 3 * time.Second}
	r, err := client.Get(url)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&respStruct)
}

func mockSearch(keywords string, respStruct *models.ApiGifs) error {
	respStruct.Data = make([]models.ApiGif, 7, 7)

	for i, _ := range respStruct.Data {
		respStruct.Data[i] = models.ApiGif{
			EmbedUrl: fmt.Sprintf("http://giphy.fake/asdf%v", i),
			Id:       strconv.Itoa(i),
		}
	}

	return nil
}
