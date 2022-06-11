package app

import (
	"InsitTestTask/model"
	"encoding/json"
	"github.com/gomarkdown/markdown"
	"io"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	md := []byte("## markdown document")
	output := markdown.ToHTML(md, nil, nil)
	io.WriteString(w, string(output))
}
func Login(w http.ResponseWriter, r *http.Request) {
	var login model.Auth
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&login)

	}
}
