package main

import (
	"GiphyApi/controllers"

	"github.com/astaxie/beego"
)

func main() {
	//"search/keyword/asdf" => { data: [...], next: 2 }
	//"search/keyword/asdf?page=2" => { data: [...], next: 3 }
	beego.SetLogger("file", `{"filename": "logs/development.log"}`)

	beego.Router("/search/:keywords", &controllers.GifCtrl{}, "get:Index")
	beego.Run()
}
