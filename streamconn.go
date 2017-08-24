package bfapi

import (
	"bufio"
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"net"
	"regexp"
	"sync"
	"time"

	"github.com/mailru/easyjson"
)

//
type StreamHandler interface {
	OnConnect(ConnectionMessage, interface{})
	OnStatus(StatusMessage, interface{})
	OnChange(ChangeMessage, interface{})
	OnClose(interface{})
}

//
type Stream struct {
	conn   net.Conn
	closer sync.Once

	// for scheduled closes
	tiMut sync.Mutex
	tiDur time.Duration
	timer *time.Timer

	chanOut chan interface{}
	handler StreamHandler
	context interface{}
}

//
func NewStream(hdlr StreamHandler, cntx interface{}) (*Stream, error) {
	var c *tls.Conn
	var err error
	var sc *Stream

	c, err = tls.Dial("tcp", streamHost, &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		InsecureSkipVerify: true,
		Renegotiation:      tls.RenegotiateFreelyAsClient,
		Rand:               rand.Reader,
	})

	if err != nil {
		return nil, err
	}

	sc = &Stream{
		conn:    c,
		chanOut: make(chan interface{}, 20),
		handler: hdlr,
		context: cntx,
	}

	// send login packets
	sc.chanOut <- AuthMessage{
		Op:      Auth,
		ID:      0,
		Session: token,
		AppKey:  appKey,
	}

	return sc, err
}

//
func (sc *Stream) Send(msg interface{}) bool {
	select {
	case sc.chanOut <- msg:
		return true
	default:
		return false
	}
}

//
func (sc *Stream) ListenWriteOut() {
	var err error
	var enc *json.Encoder = json.NewEncoder(sc.conn)

	for m := range sc.chanOut {
		if err = enc.Encode(m); err != nil {
			sc.Close()
			break
		}
	}

	sc.Close()
}

//
func (sc *Stream) ListenReadIn() {
	var err error
	var in *bufio.Reader = bufio.NewReader(sc.conn)
	var re *regexp.Regexp = regexp.MustCompile(`"op"\s*:\s*"([a-zA-Z]+)\s*"`)

	for {
		var jsonBytes []byte
		if jsonBytes, err = in.ReadBytes('\n'); err != nil {
			sc.Close()
			return
		}

		var op string = string(re.FindSubmatch(jsonBytes)[1])

		switch op {
		case Conn:
			var cm ConnectionMessage
			err = easyjson.Unmarshal([]byte(jsonBytes), &cm)
			if err == nil {
				sc.handler.OnConnect(cm, sc.context)
			}

		case Status:
			var sm StatusMessage
			err = easyjson.Unmarshal([]byte(jsonBytes), &sm)
			if err == nil {
				sc.handler.OnStatus(sm, sc.context)
			}

		case Mcm:
			var mc ChangeMessage
			err = easyjson.Unmarshal([]byte(jsonBytes), &mc)
			if err == nil {
				sc.handler.OnChange(mc, sc.context)
			}

		case Ocm:
			var oc ChangeMessage
			err = easyjson.Unmarshal([]byte(jsonBytes), &oc)
			if err == nil {
				sc.handler.OnChange(oc, sc.context)
			}
		}
	}
}

//
func (sc *Stream) Close() {
	sc.closer.Do(func() {
		sc.conn.Close()
		close(sc.chanOut)
		sc.handler.OnClose(sc.context)
	})
}

//
func (sc *Stream) ScheduleClose(in time.Duration) {
	sc.tiMut.Lock()
	if sc.timer == nil {
		sc.tiDur = in
		sc.timer = time.AfterFunc(in, sc.Close)
	}
	sc.tiMut.Unlock()
}

//
func (sc *Stream) ResetClose() {
	sc.tiMut.Lock()
	if sc.timer != nil && sc.timer.Stop() {
		sc.timer.Reset(sc.tiDur)
	}
	sc.tiMut.Unlock()
}
