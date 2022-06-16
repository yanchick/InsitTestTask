package main

import (
	"github.com/yanchick/InsitFrontendTest/app"
	"github.com/tidwall/buntdb"
	"net/http"
)

func main() {
	db, _ := buntdb.Open(":memory:")

	http.HandleFunc("/info", app.Info(db))
	http.HandleFunc("/login", app.Login(db))
	http.HandleFunc("/task", app.Task(db))
	http.ListenAndServe(":1080", nil)
}
