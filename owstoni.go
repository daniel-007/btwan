package btwan

import (
	"flag"
	"io"
	"math/rand"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

var (
	workdir string
	rpcAddr string
	logfile string
	ow      *owstoni
)

func init() {
	flag.StringVar(&workdir, "workdir", "/var/lib/owstoni", "work directory")
	flag.StringVar(&logfile, "logfile", "stdout", "log file")
	flag.StringVar(&rpcAddr, "rpcAddr", ":7700", "rpc address")
	snow, err := NewNode(int64(rand.Intn(1023)))
	if err != nil {
		panic(err)
	}
	snowflake = snow
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

	if err := _initIndexer(); err != nil {
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
	return s.Serve(l)
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
	return getMetadata(ih.ID)

}
func (o *owstoni) Index(_ context.Context, m *MetadataInfo) (*Void, error) {
	//info("request.Index", m)
	if err := m.save(); err != nil {
		return nil, err
	}
	_indexChan <- m
	return &Void{}, nil
}
func (o *owstoni) Search(_ context.Context, req *SearchReq) (*SearchResp, error) {

	resp := searchIndex(req.Q, int(req.Offset), int(req.Limit))
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
	return &result, nil
}

var snowflake *Node

//GenrateID ....
func GenrateID() int64 {
	return snowflake.Generate().Int64()
}
