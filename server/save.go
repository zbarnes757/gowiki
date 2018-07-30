package server

import (
	"gowiki/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) saveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	body := r.FormValue("body")
	page := models.Page{Title: vars["title"], Body: body}

	if err := s.db.Model(&page).Where("title = ?", vars["title"]).Update(page).Error; err != nil {
		http.Error(w, err.Error(), 400)
	} else {
		http.Redirect(w, r, "/view/"+page.Title, http.StatusFound)
	}
}
