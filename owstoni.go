package btwan

import (
	"flag"
	"io"
	"net"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

var (
	workdir string
	rpcAddr string
	logfile string
	ow      *owstoni
	laddr   string
)

func init() {
	flag.StringVar(&workdir, "workdir", "/var/lib/owstoni", "work directory")
	flag.StringVar(&logfile, "logfile", "stdout", "log file")
	flag.StringVar(&rpcAddr, "rpcAddr", ":7700", "rpc address")
	flag.StringVar(&laddr, "laddr", ":8800", "http address")
	flag.Parse()
}

func _init() error {
	if !flag.Parsed() {
		flag.Parse()
	}
	if err := initLog(); err != nil {
		return err
	}
	if err := initDB(); err != nil {
		return err
	}

	if err := initIndex(); err != nil {
		return err
	}
	return nil
}

//ListenAndServe ....
func ListenAndServe() error {
	err := _init()
	if err != nil {
		fatal(err)
		return err
	}
	l, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		fatal(err)
		return err
	}
	defer l.Close()
	s := grpc.NewServer()
	ow = &owstoni{ch: make(chan *Event, 100)}
	RegisterOwstoniServiceServer(s, ow)
	info("Owstoni Listen at", l.Addr().String())
	go s.Serve(l)
	info("Owstoni Http Listen at", laddr)
	return serveHTTP(laddr)
}

type owstoni struct {
	ch chan *Event
}

func (o *owstoni) sendEvent(e *Event) {
	o.ch <- e
}
func (o *owstoni) Send(stream OwstoniService_SendServer) error {
	for {
		event, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		info(event)
	}
}
func (o *owstoni) Recv(_ *Void, stream OwstoniService_RecvServer) error {
	for {
		e := <-o.ch
		if e != nil {
			if err := stream.Send(e); err != nil {
				return err
			}
		}
	}
}
func newEvent(ttype string) *Event {
	return &Event{Type: ttype, Attributes: map[string]string{}}
}
func (o *owstoni) SendInfoHash(_ context.Context, ih *InfoHash) (*Void, error) {
	event := newEvent("req.collect")
	event.Attributes["infohash"] = ih.Ih
	o.ch <- event
	return &Void{}, nil
}

func (o *owstoni) GetMetadataInfo(_ context.Context, ih *InfoHash) (*MetadataInfo, error) {
	return getMetadata(strconv.FormatUint(ih.ID, 10))

}
func (o *owstoni) Index(_ context.Context, m *MetadataInfo) (*Void, error) {
	//info("request.Index", m)
	if strings.ContainsAny(m.Name, "�") {
		return &Void{}, nil
	}
	if err := m.save(); err != nil {
		return nil, err
	}
	_indexChan <- m
	return &Void{}, nil
}
func (o *owstoni) Search(_ context.Context, req *SearchReq) (*SearchResp, error) {
	resp, err := bleveSearch(req.Q, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	ids := []string{}
	for _, item := range resp.Hits {
		info(item.HitNumber, item.ID, item.Score, item.Sort, item.Fields)
		info(item.String())
		ids = append(ids, item.ID)
	}
	result := SearchResp{}
	result.Request = req
	result.TotalCount = uint32(resp.Total)
	info(req, ids)
	ms, err := findMetadata(ids)
	if err != nil {
		fatal(err)
	}
	result.Metainfos = ms
	return &result, nil
}
