package main

import (
	"gowiki/database"
	"gowiki/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	db := database.SetupDB()
	defer db.Close()

	s := server.NewServer(db)
	s.RegisterRouter(r)

	n := negroni.Classic()
	n.UseHandler(r)

	log.Fatal(http.ListenAndServe(":3000", n))
}
