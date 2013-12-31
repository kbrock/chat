package main

import (
	"flag"
	"go/build"
	"net/http"
	"log"
	"github.com/kbrock/chat"
)

var (
	addr    = flag.String("addr", ":8080", "http service address")
	webroot = flag.String("root", defaultRoot(), "path to web root")
)

// thanks gary.burd.info/go-websocket-chat
func defaultRoot() string {
	p, err := build.Default.Import("github.com/kbrock/chat", "", build.FindOnly)
  if err == nil {
    return p.Dir+"/chat/app"
   } else {
     return "./app"
   }
}

func main() {
	flag.Parse()
	server := chat.NewServer()
	go server.Route()

	http.Handle("/entry", chat.NewClientHandler(server))
	http.Handle("/", http.FileServer(http.Dir(*webroot)))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
