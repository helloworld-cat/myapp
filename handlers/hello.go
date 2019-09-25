package handlers

import (
	"github.com/pagedegeek/basicuuidgenerator"
	"github.com/pagedegeek/myapp/core"
	"log"
	"net/http"
	"os"
)

type hello struct {
	instanceID string
}

func NewHelloHandler() *hello {
	uuid, _ := basicuuidgenerator.NewV4()
	return &hello{instanceID: uuid}
}

func (h *hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Printf("%s -- [%s] %s %s", h.instanceID, req.Method, req.RemoteAddr, req.URL.Path)
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(http.StatusOK)
	name := os.Getenv("FOO")
	msg := core.SayHello(name)
	rw.Write([]byte(msg))
}
