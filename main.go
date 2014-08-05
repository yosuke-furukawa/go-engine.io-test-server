package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/googollee/go-engine.io"
)

func main() {
	server, err := engineio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			conn, _ := server.Accept()
			go func() {
				defer conn.Close()
				for {
					t, r, _ := conn.NextReader()
					if r != nil {
						b, _ := ioutil.ReadAll(r)
						r.Close()
						if t == engineio.MessageText {
							log.Println(t, string(b))
						} else {
							log.Println(t, hex.EncodeToString(b))
						}
            w, _ := conn.NextWriter(t)
              w.Write([]byte("pong"))
              w.Close()
					}
				}
			}()
		}
	}()

	http.Handle("/", server)
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
