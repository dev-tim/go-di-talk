package server

import (
	"fmt"
	"gh.dev-tim/di-talk/config"
	"gh.dev-tim/di-talk/internal/storage"
)

type Server struct {
	Storage storage.Storer
	Config  *config.Config
}

func NewServer(s storage.Storer, c *config.Config) (*Server, func()) {
	return &Server{Storage: s, Config: c}, func() {
		fmt.Println("Server cleanup function called.")
	}
}

func (s *Server) Start() {
	fmt.Println("Server started on port:", s.Config.ServerPort)
	message := s.Storage.RetrieveMessage()
	fmt.Printf("Retrieved message from storage: %s\n", message)
}
