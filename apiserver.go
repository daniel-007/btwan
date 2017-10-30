package btwan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var router = httprouter.New()

func serveHTTP(laddr string) error {
	return http.ListenAndServe(laddr, cors.Default().Handler(router))
}
func init() {
	router.PanicHandler = panicHandler
	router.GET("/search/:q", search)
	router.GET("/search", search)
	router.GET("/suggest/:q", suggest)
	router.GET("/suggest", suggest)
	router.GET("/info/:id", infohash)
	router.GET("/info", infohash)
}

// @Private reason
func panicHandler(w http.ResponseWriter, _ *http.Request, err interface{}) {
	log.Println(err)
	renderError(w, "Internal Server Error", 500)
}
func infohash(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	q := p.ByName("id")
	if q == "" {
		q = r.FormValue("id")
	}
	meta, err := getMetadata(q)
	if err != nil {
		fatal(err)
		renderError(w, "获取信息错误", 422)
		return
	}
	renderJSON(w, meta, 200)
}
func suggest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	q := p.ByName("q")
	if q == "" {
		q = r.FormValue("q")
	}
	list := prefixSuggest(q)
	renderJSON(w, list, 200)
}

func search(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	q := p.ByName("q")
	if q == "" {
		q = r.FormValue("q")
	}
	offset := r.FormValue("offset")
	limit := r.FormValue("limit")
	off, _ := strconv.Atoi(offset)
	lim, _ := strconv.Atoi(limit)
	if off < 0 {
		off = 0
	}
	if lim <= 0 || lim > 100 {
		lim = 10
	}
	req := &SearchReq{Q: q, Offset: uint32(off), Limit: uint32(lim)}
	resp, err := bleveSearch(req.Q, int(req.Offset), int(req.Limit))
	if err != nil {
		fatal(err)
		renderError(w, "检索错误", 422)
		return
	}
	ids := []string{}
	for _, item := range resp.Hits {
		//info(item.HitNumber, item.ID, item.Score, item.Fragments, item.Sort, item.Fields)
		ids = append(ids, item.ID)
	}
	result := SearchResp{}
	result.Request = req
	result.TotalCount = uint32(resp.Total)
	info(req, ids)
	ms, err := findMetadata(ids)
	if err != nil {
		fatal(err)
		renderError(w, "查询错误", 422)
		return
	}

	result.Metainfos = filterMetaInfo(ms)
	renderJSON(w, &result, 200)
}

func filterMetaInfo(ms []*MetadataInfo) []*MetadataInfo {
	list := make([]*MetadataInfo, 0)
	for _, m := range ms {
		l := uint64(0)
		files := make([]*FileInfo, 0)
		for i, f := range m.Files {
			l = l + f.Length
			if i <= 2 {
				files = append(files, f)
			}
		}
		m.Length = l
		m.FileLength = uint64(len(m.Files))
		m.Files = files
		list = append(list, m)
	}
	return list
}

func renderJSON(w http.ResponseWriter, ret interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	b, _ := json.Marshal(ret)
	w.Write(b)
}

func renderError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", msg)))
}
