package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/zeeraw/eurofurence-reg/database"

	"github.com/gorilla/mux"
)

func main() {
	b, err := ioutil.ReadFile("database.json")
	if err != nil {
		log.Fatalln(err)
	}
	db := &database.Database{}
	json.Unmarshal(b, &db)

	r := mux.NewRouter()
	r.HandleFunc("/register/{id}/{accessToken}", registerHandler(db))
	s := &http.Server{
		Handler: r,
		Addr:    ":7373",
	}
	s.ListenAndServe()
}

func registerHandler(db *database.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		for _, r := range db.Registrants {
			id := vars["id"]
			accessToken := vars["accessToken"]
			if r.ID == id && r.AccessToken == accessToken {
				f, err := os.Open(path.Join("static", id+".html"))
				defer f.Close()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				io.Copy(w, f)
				return
			}
		}
		http.Error(w, "cant find registration", http.StatusNotFound)
	}
}
