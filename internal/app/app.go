package app

import "Messenger/Messenger/internal/transport"

const httpPort = "8080"

func Run() {
	httpServer := transport.New(httpPort)
	httpServer.StartHttpServer()
	print("OK")
}
