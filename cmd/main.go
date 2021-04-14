package main

import (
	"github.com/serhiy-v/todo-app.git/config"
	"github.com/serhiy-v/todo-app.git/server"

	"log"
)

func main()  {
	cfg,err := config.Load("config/config.json")
	if err != nil{
		log.Fatal(nil)
	}
	s := server.New(cfg)
	s.Run()
}
