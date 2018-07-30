package main

import (
	"gowiki/database"
	"gowiki/server"
	"log"
	"net/http"
	"os"

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

	port := getPort()

	log.Fatal(http.ListenAndServe(":"+port, n))
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}

	return "3000"
}
