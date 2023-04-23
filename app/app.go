package app

import (
	"encoding/json"
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gorilla/sessions"
	"github.com/tidwall/buntdb"
	"github.com/yanchick/InsitFrontendTest/model"
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
			session, _ := store.Get(r, "insit")
			session.Values["authenticated"] = true
			session.Values["username"] = login.Login
			session.Save(r, w)
			bolB, _ := json.Marshal(true)
			fmt.Fprintf(w, string(bolB))
		}
		if r.Method == "OPTIONS" {
			fmt.Fprintf(w, "")
		}
	}

	return http.HandlerFunc(fn)
}

func Task(db *buntdb.DB) http.HandlerFunc {
	//var tasks []model.Task
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "insit")
		var isAuth = !session.IsNew

		if isAuth {
			var name = session.Values["username"].(string)

			switch r.Method {
			case "GET":
				_ = db.View(func(tx *buntdb.Tx) error {
					val, err := tx.Get(name)
					if err != nil {
						return err
					}
					fmt.Fprintf(w, val)
					return nil
				})
			case "PUT":
				var tasks []model.Task
				var task model.Task
				bodyBytes, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Fatal(err)
				}
				_ = json.Unmarshal(bodyBytes, &task)
				err = db.View(func(tx *buntdb.Tx) error {
					val, err := tx.Get(name)
					err = json.Unmarshal([]byte(val), &tasks)
					if err != nil {
						return err
					}
					return nil
				})

				if err != nil {
					log.Fatal(err)
				}
				for i, s := range tasks {
					if s.Description == task.Description {
						tasks[i].State = task.State
					}
				}
				err = db.Update(func(tx *buntdb.Tx) error {
					b, _ := json.Marshal(tasks)
					tx.Set(name, string(b), nil)
					return nil
				})
				if err != nil {
					log.Fatal(err)
				}

			case "OPTIONS":
				fmt.Fprintf(w, "")

			case "POST":
				err := db.Update(func(tx *buntdb.Tx) error {

					bodyBytes, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Fatal(err)
					}
					bodyString := string(bodyBytes)
					tx.Set(name, bodyString, nil)
					return nil

				})
				if err != nil {
					log.Fatal(err)
				}

			}
		} else {
			w.WriteHeader(http.StatusForbidden)

		}
	}
	return http.HandlerFunc(fn)
}
