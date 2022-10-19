package webssh

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func SSHHandle(wsConn *websocket.Conn) {
	client, err := newSSHClient(
		WithHostAddr("hostAddress"),
		WithKeyValue("pubkey"),
		WithUser("intel"),
		WithTimeOut(time.Second),
	)
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
