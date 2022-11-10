package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	host = flag.String("h", "localhost", "Host to listen on")
	port = flag.Int("p", 3090, "Port to listen on")
)

func main() {
	flag.Parse()

	//Crea una conexion hacia el servidor
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected")
	done := make(chan struct{})

	//De manera concurrente se escucha la conexion y se muestran en consola
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done

}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
