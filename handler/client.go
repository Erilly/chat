package handler

import (
	"golang.org/x/net/websocket"
	"sync"

)

type Client struct {
	ID     int
	Conn   *websocket.Conn
	Queue  chan []byte
	WrMux  sync.RWMutex
	StMux  sync.RWMutex
	Status int8
}

func (this *Client) Run() {
	defer this.Conn.Close()
}

func (this *Client) SetStatus(status int8) {
	defer this.StMux.Unlock()
	this.StMux.Lock()
	this.Status = status
}

func (this *Client) GetStatus() int8 {
	defer this.StMux.RUnlock()
	this.StMux.RLock()
	return this.Status
}