package handlers

import (
	"log"
	"net/http"
)

type goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *goodbye {
	return &goodbye{l}
}

func (g *goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Bye"))
}
