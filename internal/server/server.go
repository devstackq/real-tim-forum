package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/devstackq/real-time-forum/internal/repository"
	"github.com/devstackq/real-time-forum/internal/service"
	"github.com/devstackq/real-time-forum/internal/handler"
)

func NewServer(conf *Config) *Server {
	db, err := repository.CreateDB(conf.DbDriver, conf.DbPath)
	if err != nil {
		log.Println(err)
	}

	repos := repository.NewRepository(db)
	fmt.Println(repos, "repo, create Db ?")
	services := service.NewService(repos)
	fmt.Println(services, "prepare services")
	
	handler := handler.NewHandler(services)

	port := os.Getenv("PORT")
	if port == "" {
		port = conf.Port
	}

	s := &Server{
		http: &http.Server{
			Addr:         port,
			Handler:      handler.InitRouter(),
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		},
	}
	return s
}
run server

func Run() {

}
