package webssh

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SSHHandle(w http.ResponseWriter, r *http.Request, opt ...Option) {
	sec := r.Header.Get("Sec-WebSocket-Protocol") // get sn
	log.Println("sec:", sec)
	h := http.Header{}
	if sec != "" {
		h.Add("Sec-WebSocket-Protocol", sec)
	} else {
		h = nil
	}
	wsConn, err := upgrader.Upgrade(w, r, h)
	if err != nil {
		return
	}
	defer wsConn.Close()
	client, err := newSSHClient(opt...)
	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte(err.Error()), time.Now().Add(time.Second))
		return
	}
	defer client.Close()

	turn, err := newTurn(wsConn, client)
	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte(err.Error()), time.Now().Add(time.Second))
		return
	}
	defer turn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := turn.LoopRead(ctx)
		if err != nil {
			log.Printf("%#v", err)
		}
	}()
	go func() {
		defer wg.Done()
		err := turn.SessionWait()
		if err != nil {
			log.Printf("%#v", err)
		}
		cancel()
	}()
	wg.Wait()
}
