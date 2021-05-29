package main

import (
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
	//write conf struct
	//create newServer
}
