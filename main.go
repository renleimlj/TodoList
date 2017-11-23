package main

import (
	"net/http"
  	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"fmt"
)

func initRoute(m *martini.ClassicMartini) {
	m.Use(render.Renderer())
  	m.Get("/", func(r render.Render) {
    	r.HTML(200, "index", nil)
	})
	m.Post("/add", func(r render.Render, req *http.Request) {
		r.JSON(200, map[string]interface{}{
			"issues" : req.FormValue("issue"),
		})
	})
	m.NotFound(func(r render.Render, req *http.Request)  {
		fmt.Println("[martini] Page Not Found[500]")
		r.HTML(500, "notfound", req.URL.Path)
	})
}

func main() {
  	m := martini.Classic()
	initRoute(m)
	m.Use(martini.Static("assets"))
	m.Run()
}