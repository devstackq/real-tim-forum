package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/devstackq/real-time-forum/internal/handler"
	"github.com/devstackq/real-time-forum/internal/repository"
	"github.com/devstackq/real-time-forum/internal/service"
)

type Server struct {
	http *http.Server
}

func NewServer(conf *Config) *Server {
	db, err := repository.CreateDB(conf.DbDriver, conf.DbPath)
	if err != nil {
		log.Println(err)
	}
	//prepare config, dependency , then run custom  server, DI
	//chain interface relation between layer -> repos->services->handlers
	//outer layer connect -> inner - with interfaces, then realize interfaces

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	port := os.Getenv("PORT")
	if port == "" {
		port = conf.Port
	}

	//custom server
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

func (s *Server) Run() error {
	log.Println("start server in port: ", s.http.Addr)
	return s.http.ListenAndServe()
}
