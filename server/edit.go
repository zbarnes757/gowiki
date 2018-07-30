package server

import (
	"gowiki/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) editHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var page models.Page
	s.db.First(&page, "title = ?", vars["title"])

	if (models.Page{}) == page {
		page.Title = vars["title"]
		s.db.Create(&page)
	}

	renderTemplate(w, "edit", &page)
}
