package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	WsMsgCodePing      = 10001
	WsMsgCodePong      = 10002
	WsMsgCodeClose     = 10003
	WsMsgCodeBroadcast = 10004
	WsMsgCodeSelfInfo  = 10005
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	for i := 1; i < 10000; i++ {
		time.Sleep(50 * time.Millisecond)
		go NewClient()
	}
	go polling()
	wg.Wait()
}

var addr = flag.String("addr", "localhost:2303", "http service address")

func NewClient() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.BinaryMessage, MsgBuffer(WsMsgCodePing, []byte("Test Message"+t.String())))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func MsgBuffer(msgCode int, body []byte) (buf []byte) {
	var (
		buyCode []byte
		buffer  bytes.Buffer
	)
	buyCode = Int2Bytes(msgCode)
	buffer.Write(buyCode)
	if body != nil {
		buffer.Write(body)
	}
	buf = buffer.Bytes()
	return
}

func Int2Bytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func polling() {
	ticker := time.NewTicker(time.Minute * 30)
	for {
		select {
		case <-ticker.C:
			for i := 1; i < 1000; i++ {
				time.Sleep(50 * time.Millisecond)
				go NewClient()
			}
		}
	}
}
