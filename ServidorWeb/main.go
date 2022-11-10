package main

func main() {

	server := NewServer(":5000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMiddleWare(HandleHome, CheckAuth(), Logging()))
	server.Handle("POST", "/create", HandlePost)
	server.Listen()

}
