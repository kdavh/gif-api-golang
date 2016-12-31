package models

type ApiGif struct {
	Id       string `json:"id"`
	EmbedUrl string `json:"embed_url"`
}

type ApiGifs struct {
	Data []ApiGif `json:"data"`
}
