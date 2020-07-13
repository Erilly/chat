package handler

import (
	"net/http"
	"encoding/json"
	"golang.org/x/net/websocket"
	"fmt"
)

func Debug(w http.ResponseWriter, r *http.Request) {
	result := ResponseResult{0, "Debug success!"}

	message, _ := json.Marshal(result)

	w.Write(message)
}

func StdWs(conn *websocket.Conn) {
	var str string
	var err error

	go func() {
		for {
			select {
			case ch := <-requestChan:
				fmt.Println(ch.Res)
				err := websocket.Message.Send(ch.Conn, ch.Res)
				fmt.Println(err)
			}
		}
	}()

	for {
		err = websocket.Message.Receive(conn, &str)
		if err != nil {
			continue
		}
		requestChan <- &clients{conn, str}
	}
}
