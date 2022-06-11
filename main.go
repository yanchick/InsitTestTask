package main

import (
	"InsitTestTask/app"
	"net/http"
)

func main() {
	http.HandleFunc("/info", app.Info)
	http.HandleFunc("/login", app.Login)
	http.ListenAndServe(":1080", nil)
}
