package websockets

import "github.com/gorilla/websocket"

type Client struct {
	hub      *Hub            // El hub al que pertenece el cliente
	id       string          // El id del cliente
	socket   *websocket.Conn // La conexión websocket
	outbound chan []byte     // Canal para enviar mensajes al cliente
}

func NewClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		hub:      hub,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

func (c *Client) Write() {

	// De manera indefinida, se leen los mensajes del canal outbound
	for {
		select {
		case message, ok := <-c.outbound:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
