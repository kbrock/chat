package chat

import (
	"log"
)

// Chat server.
type Server struct {
	clients     map[*Client]bool
	addClient    chan *Client
	removeClient chan *Client
	sendAll      chan *Message
	messages     []*Message
}

// Create new chat server.
func NewServer() *Server {
  clients     := make(map[*Client]bool)
	addClient    := make(chan *Client)
	removeClient := make(chan *Client)
	sendAll      := make(chan *Message)
	messages     := make([]*Message, 0)
	return &Server{clients, addClient, removeClient, sendAll, messages}
}

func (self *Server) AddClient() chan<- *Client {
	return (chan <- *Client)(self.addClient)
}

func (self *Server) RemoveClient() chan<- *Client {
	return (chan <- *Client)(self.removeClient)
}

// send a message to all clients
func (self *Server) SendAll() chan<-*Message {
	return (chan <- *Message)(self.sendAll)
}

// get the list of all messages
func (self *Server) Messages() []*Message {
	msgs := make([]*Message, len(self.messages))
	copy(msgs, self.messages)
	return msgs
}

// Listen and serve.
// It serves client connection and broadcast request.
func (self *Server) Route() {
	for {
		select {

		// Add new a client
		case c := <-self.addClient:
			log.Println("Added new client")
		  self.clients[c] = true
			for _, msg := range self.messages {
				c.Write() <- msg
			}
			log.Println("Now", len(self.clients), "clients connected.")

		// remove a client
		case c := <-self.removeClient:
			log.Println("Remove client")
 			delete(self.clients, c)

		// broadcast message for all clients
		case msg := <-self.sendAll:
			log.Println("Send all:", msg)
			self.messages = append(self.messages, msg)
			for c := range self.clients {
				c.Write() <- msg
			}
		}
	}
}
