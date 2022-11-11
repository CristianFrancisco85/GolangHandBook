package websockets

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Actualiza la conexión del cliente a una que soporte WebSockets
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub struct {
	clients    []*Client    // Lista de clientes conectados
	register   chan *Client // Canal para registrar nuevos clientes
	unregister chan *Client // Canal para desconectar clientes
	mutex      *sync.Mutex  // Mutex para proteger la lista de clientes
}

func NewHub() *Hub {
	return &Hub{
		clients:    []*Client{},
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (h *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	// Se actualiza la conexión a una que soporte WebSockets
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	// Se crea un nuevo cliente y se registra en el canal de registro del hub
	client := NewClient(h, socket)
	h.register <- client
	// Se ejecuta como goroutine la función que lee los mensajes del entrantes del cliente
	go client.Write()
}

// Se inicia la escucha de los canales de registro y desconexión de clientes del hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.OnConnect(client)
		case client := <-h.unregister:
			h.OnDisconnect(client)
		}
	}
}

// Cuando un cliente se conecta, se agrega a la lista de clientes del hub
// Se utiliza el mutex para proteger la lista de lectura y escritura concurrente
func (h *Hub) OnConnect(client *Client) {
	log.Println("Client connected: ", client.socket.RemoteAddr())
	h.mutex.Lock()
	defer h.mutex.Unlock()
	// Se establece el id del cliente como la dirección IP del socket
	client.id = client.socket.RemoteAddr().String()
	// Se agrega el cliente a la lista de clientes del hub
	h.clients = append(h.clients, client)
}

// Cuando un cliente se desconecta, se agrega a la lista de clientes del hub
// Se utiliza el mutex para proteger la lista de lectura y escritura concurrente
func (h *Hub) OnDisconnect(client *Client) {
	log.Println("Client disconnected: ", client.socket.RemoteAddr())
	client.socket.Close()

	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Se busca el cliente en la lista de clientes del hub
	// TODO: No seria necesario buscar el cliente, si se utiliza un map
	for i, c := range h.clients {
		if c.id == client.id {
			h.clients = append(h.clients[:i], h.clients[i+1:]...)
			break
		}
	}
}

// Se envía un mensaje a todos los clientes del hub
func (h *Hub) Broadcast(message interface{}, ignore *Client) {

	data, _ := json.Marshal(message)

	h.mutex.Lock()
	defer h.mutex.Unlock()

	// Se recorre la lista de clientes del hub
	for _, client := range h.clients {
		if client != ignore {
			client.outbound <- data
		}
	}
}
