package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"time"
)

func main() {
	beego.Router("/search/:keywords", &gifController{}, "get:Index")
	beego.Run()
}

type gif struct {
	GifId string `json:"gif_id"`
	Url   string `json:"url"`
}

type gifs struct {
	Data []gif `json:"data"`
}

type gifController struct {
	beego.Controller
}

func (c *gifController) Index() {
	apiResp := new(apiGifs)

	err := getFromGiphyAPI(c.Ctx.Input.Param(":keywords"), apiResp)

	respStruct := gifs{
		Data: make([]gif, 0, 0),
	}

	if err != nil || len(apiResp.Data) >= 5 {
		respStruct.Data = make([]gif, 5, 5)
		for i := 0; i < 5; i++ {
			respStruct.Data[i] = gif{
				GifId: apiResp.Data[i].Id,
				Url:   apiResp.Data[i].Url,
			}
		}
	}

	c.Data["json"] = respStruct
	c.ServeJSON()
}

type apiGif struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type apiGifs struct {
	Data []apiGif `json:"data"`
}

func getFromGiphyAPI(keywords string, target interface{}) error {
	GIPHY_API_BETA_KEY := "dc6zaTOxFJmzC"

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

	return json.NewDecoder(r.Body).Decode(&target)
}
