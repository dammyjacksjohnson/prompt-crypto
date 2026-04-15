package hub

import (
	"sync"
)

// Hub maintains the set of active connections and broadcasts messages to the clients.
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	sync.RWMutex
}

// Client represents a single connection to the Hub.
type Client struct {
	hub  *Hub
	ch   chan []byte
}

// NewHub initializes a new Hub.
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run starts the Hub to listen for incoming connections and broadcasting messages.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			 h.Unlock()
		case client := <-h.unregister:
			h.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
			close(client.ch)
			}
			h.Unlock()
		case message := <-h.broadcast:
			h.RLock()
			for client := range h.clients {
				client.ch <- message
			}
			h.RUnlock()
		}
	}
}