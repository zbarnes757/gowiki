package server

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server is the base class for a server
type Server struct {
	db *gorm.DB
}

// NewServer instantiates a new server object with a database
func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

// RegisterRouter applies all the routes to a server
func (s *Server) RegisterRouter(r *mux.Router) {
	r.HandleFunc("/view/{title}", s.viewHandler)
	r.HandleFunc("/edit/{title}", s.editHandler)
	r.HandleFunc("/save/{title}", s.saveHandler)
	r.HandleFunc("/", s.indexHandler)
}
