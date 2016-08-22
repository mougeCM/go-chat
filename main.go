package main

import (
	"net/http"
	"gopkg.in/macaron.v1"
	"time"
	"fmt"
)

// 进入模版页面
func serveHome(ctx *macaron.Context) {
	w := ctx.Resp
	if ctx.Req.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if ctx.Req.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println(ctx.Req.Host)
	ctx.Data["host"] = ctx.Req.Host
	ctx.HTML(200, "home")
}

func route() {
	m := macaron.Classic()
	// Render
	m.Use(macaron.Static("templates/statics", macaron.StaticOptions{
		Prefix:      "statics",
		SkipLogging: true,
		IndexFile:   "inex.html",
		Expires: func() string {
			return time.Now().Add(24 * 60 * time.Minute).UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
		},
	}))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		IndentJSON: false,
		Directory:  "templates",
	}))
	// 健康检查
	m.Head("/", func(ctx *macaron.Context) {
		ctx.JSON(200, "ok")
	})

	m.Any("/", serveHome)
	m.Any("ws", serveWs)

	m.Run()
}



func main() {
	go h.run()
	route()
}
