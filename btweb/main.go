package main

import (
	"flag"

	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

var dev bool
var indexer string
var appdir string

func init() {
	flag.BoolVar(&dev, "dev", true, "dev mode,this will reload the templates on each request")
	flag.StringVar(&indexer, "indexer", "api.btwan.net:7700", "indexer server")
	flag.StringVar(&appdir, "appdir", "./views", "application templates directory")
}
func main() {
	flag.Parse()
	_init()
	iris.Config.IsDevelopment = dev
	iris.Config.Gzip = true
	iris.Config.Charset = "UTF-8"
	iris.Static("/assets", appdir+"/assets", 1)
	iris.UseTemplate(html.New()).Directory(appdir, ".html")
	iris.Listen(":9000")
}
