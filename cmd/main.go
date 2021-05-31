package main

import (
	"fmt"
	"log"

	"github.com/devstackq/real-time-forum/internal/server"
)

// create new server(config)
//start server
func main() {

	conf := server.NewConfig()

	if err := server.ReadConfig("../config/config.json", conf); err != nil {
		log.Println(err)
	}
	s := server.NewServer(conf)
	err := s.Run()
	fmt.Println(err, 123)
	if err != nil {
		log.Fatal(err)
	}
}
