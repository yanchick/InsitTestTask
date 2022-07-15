package main

import (
	"github.com/rs/cors"
	"github.com/tidwall/buntdb"
	"github.com/yanchick/InsitFrontendTest/app"
	"net/http"
)

func main() {
	db, _ := buntdb.Open(":memory:")

	mux := http.NewServeMux()
	mux.HandleFunc("/info", app.Info(db))
	mux.HandleFunc("/login", app.Login(db))
	mux.HandleFunc("/task", app.Task(db))
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}).Handler(mux)
	http.ListenAndServe(":1080", c)

}
