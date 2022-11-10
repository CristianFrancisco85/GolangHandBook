package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	mainMessages    = make(chan string)
)

var (
	host = flag.String("h", "localhost", "Host to listen on")
	port = flag.Int("p", 3090, "Port to listen on")
)

// Client1 -> Server -> HandleConnection(Client1)
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	// Se crea un canal de mensajes para el cliente
	message := make(chan string)
	//Se escucha el canal concurrentemente
	go MessageWrite(conn, message)

	//Se da la bienvenida al cliente
	clientName := conn.RemoteAddr().String()
	message <- "Welcome to the server, your name is " + clientName

	//Se notifica en el canal principal del nuevo cliente
	mainMessages <- clientName + " has joined!"
	//Se agrega el cliente al canal de clientes
	incomingClients <- message

	//Se lee la conexion para leer los mensajes entrantes del cliente
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		mainMessages <- clientName + ": " + inputMessage.Text()
	}

	//Cuando el cliente termina la conexion se
	leavingClients <- message
	mainMessages <- clientName + " has left!"
}

//Se escucha este canal para mandarle los mensajes al cliente
func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}

func Broadcast() {

	clients := make(map[Client]bool)

	for {
		// Multiplexar los mensajes
		select {
		// Nuevo Mensaje
		case msg := <-mainMessages:
			// Se envia mensaje a los clientes
			for client := range clients {
				client <- msg
			}
		// Tenemos un nuevo cliente
		case newClient := <-incomingClients:
			clients[newClient] = true
		//Un cliente se ha desconectado
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	flag.Parse()

	// Crea el server y lo escucha
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// Inicia el broadcast
	go Broadcast()

	// Escucha por conexiones
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go HandleConnection(conn)
	}
}
