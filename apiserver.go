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
}

// @Private reason
func panicHandler(w http.ResponseWriter, _ *http.Request, err interface{}) {
	log.Println(err)
	renderError(w, "Internal Server Error", 500)
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
	if lim <= 0 {
		lim = 10
	}
	req := &SearchReq{Q: q, Offset: uint32(off), Limit: uint32(lim)}
	resp := searchIndex(q, off, lim)
	result := SearchResp{}
	result.Request = req
	result.TotalCount = uint32(resp.NumDocs)
	result.Count = uint32(len(resp.Docs))
	ids := []uint64{}
	for _, doc := range resp.Docs {
		ids = append(ids, doc.DocId)
	}
	info(req, ids)
	ms, err := findMetadata(ids)
	if err != nil {
		fatal(err)
	}
	result.Metainfos = ms
	renderJSON(w, &result, 200)
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
