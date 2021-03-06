package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

type edit struct {
	ID        int
	Namespace template.HTML
	Path      string
	Title     template.HTML
	Tags      string
	Text      template.HTML
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	guestmodechek(w, r)

	if checkLogin(r) == false {
		fmt.Fprintf(w, "ERROR YOU ARE NOT LOGGED IN")
		return
	}

	u, err := url.Parse(r.URL.Path)

	checkErr(err)
	encodetpath1 := strings.Split(u.Path, "/")

	if len(encodetpath1) < 4 {
		fmt.Fprintf(w, "ERROR 404")
	} else {

		ids, err := db.Query("SELECT id,namespace,title,tags,text FROM article WHERE title=(?) AND namespace=(?)", encodetpath1[3], encodetpath1[2])
		defer ids.Close()
		checkErr(err)

		ids.Next()
		var id int
		var namespace string
		var title string
		var tags string
		var text string
		_ = ids.Scan(&id, &namespace, &title, &tags, &text)
		checkErr(err)

		if id == 0 {
			fmt.Fprintf(w, "ERROR 404")
		} else {

			NamespaceTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(namespace)))
			TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title)))
			TextTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(text)))

			edits := edit{id, NamespaceTMP, encodetpath1[2], TitleTMP, tags, TextTMP}
			templatesEdit.Execute(w, edits)
		}

	}
}
