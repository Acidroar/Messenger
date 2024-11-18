package transport

import (
	"log"
)

type HttpServer struct {
	address string
	*echo.Echo
}

func New(address string) *HttpServer {
	echoServer := echo.New()

	return &HttpServer{
		address: address,
		Echo:    echoServer,
	}
}

func (h *HttpServer) StartHttpServer() {
	err := h.Start(h.address)
	if err != nil {
		log.Fatal(err)
	}
}
