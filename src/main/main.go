package main

////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

////////////////////////////////////////////////////////////////////////////

var (
	source = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	client = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	clientConn *websocket.Conn
)

////////////////////////////////////////////////////////////////////////////

func main() {
	http.HandleFunc("/websocket/source", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("websocket source")

		src, err := source.Upgrade(w, r, nil)
		if err != nil {
			panic(err.Error())
		}

		for {
			msgType, msg, err := src.ReadMessage()
			if err != nil {
				panic(err.Error())
			}

			if clientConn == nil {
				continue
			}

			if err = clientConn.WriteMessage(msgType, msg); err != nil {
				panic(err.Error())
			}
		}
	})

	http.HandleFunc("/websocket/client", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("websocket client")

		clientConn, _ = client.Upgrade(w, r, nil) // error ignored for sake of simplicity
	})

	http.HandleFunc("/source", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/source.html")
	})

	http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/client.html")
	})

	http.ListenAndServe(":8080", nil)
}

////////////////////////////////////////////////////////////////////////////
