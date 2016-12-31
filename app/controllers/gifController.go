package controllers

import (
	"GiphyApi/models"
	"GiphyApi/services"
	"github.com/astaxie/beego"
)

type GifCtrl struct {
	beego.Controller
}

const FIRST_PAGE = 1
const PER_PAGE = 5

func (c *GifCtrl) Index() {
	page, err := c.GetInt("page")
	if err != nil || page < FIRST_PAGE {
		page = FIRST_PAGE
	}

	apiResp := new(models.ApiGifs)
	apiErr := services.GiphyClient().Search(c.Ctx.Input.Param(":keywords"), apiResp)

	respStruct := models.Gifs{
		Data: make([]models.Gif, 0, PER_PAGE),
		Next: page + 1,
	}

	if apiErr == nil {
		startIndex := PER_PAGE * (page - 1)
		endIndex := (startIndex + PER_PAGE) - 1
		if endIndex > (len(apiResp.Data) - 1) {
			endIndex = len(apiResp.Data) - 1
		}

		for i := startIndex; i <= endIndex; i++ {
			respStruct.Data = append(respStruct.Data, models.Gif{
				GifId: apiResp.Data[i].Id,
				Url:   apiResp.Data[i].EmbedUrl,
			})
		}
	}

	c.Data["json"] = respStruct
	c.ServeJSON()
}
