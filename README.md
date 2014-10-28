Experimenting in Go, WebSockets, and Angular.js

ABANDONED

working on main portions of the go chat - you can find that at [multichat](http://github.com/kbrock/multichat)
The angular code is not over there yet. That may come after a little while.

## Credit

The server portion is based upon https://github.com/tenntenn/golang-samples/tree/master/websocket/websocket-chat

Borrowed some ideas from http://gary.burd.info/go-websocket-chat :

- using map instead of a list of connections
- setting a parameter for the static root


The client is based upon http://parroty00.wordpress.com/2013/07/15/eventmachine-websocket-angularjs/


## Installation 

```bash
export GOPATH=$(pwd)

# while developing:
# mkdir -p src/github.com/kbrock
# ln -s ../../.. src/github.com/kbrock/chat
go get github.com/kbrock/chat
go install -ldflags "-X main.Build 'v1'" github.com/kbrock/chat
```

`open http://localhost:8080/`
