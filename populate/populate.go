package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/zeeraw/eurofurence-reg/database"
)

type email struct {
	Keyword    string
	Registrant database.Registrant
}

type html struct {
	EmailTemplate string
}

func main() {
	b, err := ioutil.ReadFile("database.json")
	if err != nil {
		log.Fatalln(err)
	}
	db := database.Database{}
	json.Unmarshal(b, &db)

	et := template.Must(template.ParseFiles("templates/email"))
	rt := template.Must(template.ParseFiles("templates/registration.html"))

	for _, registrant := range db.Registrants {
		e := email{
			Keyword:    db.Keyword,
			Registrant: registrant,
		}
		buf := bytes.NewBuffer(nil)
		et.Execute(buf, &e)
		h := html{
			EmailTemplate: buf.String(),
		}
		f, err := os.Create(path.Join("static", registrant.ID+".html"))
		if err != nil {
			log.Fatalln(err)
		}
		err = rt.Execute(f, &h)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
