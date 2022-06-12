package app

import (
	"InsitTestTask/model"
	"encoding/json"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gorilla/sessions"
	"github.com/tidwall/buntdb"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Info(db *buntdb.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		md := []byte("## markdown document")
		output := markdown.ToHTML(md, nil, nil)
		io.WriteString(w, string(output))
	}
	return http.HandlerFunc(fn)
}

func Login(db *buntdb.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var login model.Auth
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&login)
			session, _ := store.Get(r, "cookie-name")
			session.Values["authenticated"] = true
			session.Values["username"] = login.Login
			session.Save(r, w)
			fmt.Fprintf(w, "ok\n")
		}
	}
	return http.HandlerFunc(fn)
}

func Task(db *buntdb.DB) http.HandlerFunc {
	//var tasks []model.Task
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		var name = session.Values["username"].(string)
		switch r.Method {
		case "GET":
			fmt.Println("one")
		case "POST":
			db.Update(func(tx *buntdb.Tx) error {
				bodyBytes, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Fatal(err)
				}
				bodyString := string(bodyBytes)
				tx.Set(name, bodyString, nil)
				return nil
			})
		case "PUT":
			fmt.Println("three")
		}

	}
	return http.HandlerFunc(fn)
}
