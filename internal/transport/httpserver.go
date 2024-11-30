package transport

import (
	"github.com/labstack/echo/v4"
	"log"
)

type HttpServer struct {
	address string
	*echo.Echo
}

func New(address string) *HttpServer {
	echoServer := echo.New()
	print("OK")
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
	print("OK")
}
