package handler

import (
	"fmt"
	"time"
	"golang.org/x/net/websocket"
)


const(
	MAX_MESSAGE_SIZE = 10240 * 10
	PONG_WAIT = 360*time.Second

	CLIENT_NORMAL  int8 = iota
	CLINET_READING
	CLIENT_WRITING
	CLIENT_EXIT

)
type clients struct{
	Conn *websocket.Conn
	Res string
}

var requestChan  = make(chan *clients)

type ResponseResult struct{
	Code	int		`json:"code"`
	Msg		string	`json:"msg"`
}

func init(){

	fmt.Println("init configures")
}


