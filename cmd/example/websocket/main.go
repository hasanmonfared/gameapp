package main

import (
	"encoding/json"
	"fmt"
	"gameapp/entity"
	"github.com/labstack/gommon/log"
	"net"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}

		defer conn.Close()

		channel := make(chan string)
		go producer(r.RemoteAddr, channel)
		go writeMessage(conn, channel)

		done := make(chan bool)
		go readMessage(conn, done)

		<-done
	}))
}
func readMessage(conn net.Conn, done chan<- bool) {
	for {
		msg, opCode, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Print(err)
			done <- true
			return
		}
		var notif = entity.Notification{}
		err = json.Unmarshal(msg, &notif)
		if err != nil {
			panic(err)
		}
		fmt.Println("opCode", opCode)
		fmt.Println("notif", notif)
	}
}
func producer(remoteAddress string, channel chan<- string) {
	for {
		channel <- remoteAddress
		time.Sleep(2 * time.Second)
	}
}
func writeMessage(conn net.Conn, channel <-chan string) {
	for data := range channel {

		err := wsutil.WriteServerMessage(conn, ws.OpText, []byte(data))
		if err != nil {
			panic(err)
		}

	}
}
