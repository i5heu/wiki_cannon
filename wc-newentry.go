package main

import (
	"fmt"
	"net/http"
	"time"
)

func NewentryHandler(w http.ResponseWriter, r *http.Request) {

	t := "login: false"

	if checkLogin(r) == true {
		//if true == true {
		t = "login: true"

		newTitle := r.FormValue("Title")
		newNamepace := r.FormValue("Namespace")
		newText := r.FormValue("Text")

		if newNamepace == "" {
			newNamepace = "main"
		}
		if newTitle == "" {
			t := time.Now()
			newTitle = t.String()[:len(t.String())-21]
		}

		db.Exec("INSERT INTO article(title,namespace,text) VALUES(?,?,?)", ReplaceSpecialChars(newTitle), ReplaceSpecialChars(newNamepace), newText)

		checkErr(err)

		http.ServeFile(w, r, "./template/newentry.html")

	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}
