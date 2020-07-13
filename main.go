package main

import (
	"net/http"
	"time"
	"fmt"
	"os"
	"chat/handler"
	"golang.org/x/net/websocket"
)

func main(){

	go func(){
		adminServer := &http.Server{
			Addr:	"127.0.0.1:8080",
			ReadTimeout:	10*time.Second,
			WriteTimeout:	10*time.Second,
			MaxHeaderBytes:	1<<20,
		}

		serverMux := http.NewServeMux()

		serverMux.HandleFunc("/debug", handler.Debug)

		adminServer.Handler = serverMux

		if err := adminServer.ListenAndServe(); err!=nil{
			fmt.Fprintf(os.Stderr,"%s\n",err)
		}
	}()

	clientServer := &http.Server{
		Addr:	"127.0.0.1:8888",
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		MaxHeaderBytes:	1<<20,
	}

	serverMux := http.NewServeMux()
	serverMux.Handle("/std", websocket.Handler(handler.StdWs))

	clientServer.Handler = serverMux

	if err := clientServer.ListenAndServe(); err !=nil {
		fmt.Fprintf(os.Stderr,"%s\n",err)
	}

	//time.Sleep(time.Hour)

	//for{
	//	select{
	//		case <- clients:
	//
	//	}
	//}
}