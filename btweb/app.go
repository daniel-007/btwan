package main

import (
	"btwan"

	"github.com/kataras/iris"
	"golang.org/x/net/context"
)

var app = App{}

func init() {
	iris.Get("/", app.index)
	iris.Get("/ih/:ih", app.submitInfohash)
	iris.Get("/s", app.search)
}

//App ....
type App struct {
}

func (c App) index(ctx *iris.Context) {
	ctx.MustRender("app/index.html", map[string]interface{}{
		"title": "BT湾 - 磁力搜索",
	})
}

func (c App) submitInfohash(ctx *iris.Context) {
	ih := ctx.Param("ih")
	client.SendInfoHash(context.Background(), &btwan.InfoHash{Ih: ih})

}
func (c App) search(ctx *iris.Context) {
	var resp *btwan.SearchResp
	q := ctx.URLParam("q")
	offset, _ := ctx.URLParamInt("off")
	limit, _ := ctx.URLParamInt("limit")
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	resp, err := client.Search(context.Background(), &btwan.SearchReq{Q: q, Offset: uint32(offset), Limit: uint32(limit)})
	ctx.MustRender("app/search.html", map[string]interface{}{
		"title":  "BT湾 - 磁力搜索 - " + q,
		"q":      q,
		"off":    offset,
		"limit":  limit,
		"result": resp.Metainfos,
		"err":    err,
	})

}
