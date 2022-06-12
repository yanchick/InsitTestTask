package main

import (
	"InsitTestTask/app"
	"github.com/tidwall/buntdb"
	"net/http"
)

func main() {
	db, _ := buntdb.Open(":memory:")

	http.HandleFunc("/info", app.Info(db))
	http.HandleFunc("/login", app.Login(db))
	http.ListenAndServe(":1080", nil)
}
