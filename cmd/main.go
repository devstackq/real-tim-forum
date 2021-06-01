package main

import (
	"log"

	"github.com/devstackq/real-time-forum/internal/server"
)

func main() {

	conf := server.NewConfig()
	if err := server.ReadConfig("../config/config.json", conf); err != nil {
		log.Println(err)
	}
	s := server.NewServer(conf)
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
