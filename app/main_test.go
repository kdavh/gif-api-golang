package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/astaxie/beego"
	log "github.com/goinggo/tracelog"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGifSearch(t *testing.T) {
	r, _ := http.NewRequest("GET", "/search/keywords", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	log.Trace("testing", "gifController", "Code[%d]\n%s", w.Code, w.Body.String())

	var response gifs

	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("gif search\n", t, func() {
		Convey("status code should be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})

		Convey("There should be json data", func() {
			So(len(response.Data), ShouldBeGreaterThan, 0)
		})
	})
}
