package main

import (
	"flag"
	"go/build"
	"net/http"
	"log"
	"github.com/kbrock/chat/chat"
)

var (
	addr    = flag.String("addr", ":8080", "http service address (e.g.: localhost:8080")
	webroot = flag.String("root", defaultRoot(), "path to web root")
)

// thanks https://github.com/mattetti/go-web-api-demo
// Note:
// compile passing -ldflags "-X main.Build <build sha1>"
var Build string

// thanks gary.burd.info/go-websocket-chat
func defaultRoot() string {
	p, err := build.Default.Import("github.com/kbrock/chat", "", build.FindOnly)
  if err == nil {
    return p.Dir+"/webroot"
   } else {
     return "./webroot"
   }
}

func main() {
	flag.Parse()
	server := chat.NewServer()
	go server.Route()

  log.Println("Build", Build)
  log.Println("About to start the server on", *addr)
	http.Handle("/entry", chat.NewClientHandler(server))
	http.Handle("/", http.FileServer(http.Dir(*webroot)))
	http.ListenAndServe(*addr, nil)
}
