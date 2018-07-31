package server

import (
	"gowiki/models"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html", "index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) viewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var page models.Page
	s.db.First(&page, "title = ?", vars["title"])

	if (models.Page{}) == page {
		http.Redirect(w, r, "/edit/"+vars["title"], http.StatusFound)
	} else {
		renderTemplate(w, "view", &page)
	}
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	var pages []models.Page

	if err := s.db.Find(&pages).Error; err != nil {
		http.Error(w, err.Error(), 400)
	} else {
		err := templates.ExecuteTemplate(w, "index.html", pages)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
