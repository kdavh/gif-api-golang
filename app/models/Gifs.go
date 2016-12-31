package models

type Gif struct {
	GifId string `json:"gif_id"`
	Url   string `json:"url"`
}

type Gifs struct {
	Data []Gif `json:"data"`
	Next int   `json:"next"`
}
