package server

import (
	"github.com/serhiy-v/todo-app.git/config"
	"log"
	"net/http"
)

type Server struct {
	config config.Config
}


func New(cfg config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

func (app *Server) Run (){
	port := app.config.Port
	log.Fatal(http.ListenAndServe(port,nil))
}