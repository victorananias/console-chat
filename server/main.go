package main

var server Server

func main() {
	server.registerRoutes()
	server.start()
}
