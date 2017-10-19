package controllers

import (
	"btwan"

	"github.com/revel/revel"
	"golang.org/x/net/context"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) SubmitInfohash(ih string) revel.Result {
	client.SendInfoHash(context.Background(), &btwan.InfoHash{Ih: ih})
	return c.Render()
}
func (c App) Search(q string, offset, limit int) revel.Result {
	revel.INFO.Println(q, offset, limit)
	var resp *btwan.SearchResp

	c.RenderArgs["q"] = q
	c.RenderArgs["off"] = offset
	c.RenderArgs["limit"] = limit
	resp, err := client.Search(context.Background(), &btwan.SearchReq{Q: q, Offset: uint32(offset), Limit: uint32(limit)})
	if err != nil {
		return c.Render()
	}
	c.RenderArgs["result"] = resp.Metainfos
	return c.Render()
}
