package server

import (
	"alexdenkk/applications/internal/writer"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Writer *writer.Writer
}

func (s *Server) Run() {
	s.Engine.Run(":8080")
}

func New(w *writer.Writer) *Server {
	srv := &Server{
		Engine: gin.Default(),
		Writer: w,
	}

	srv.RegisterEndpoints()
	return srv
}
