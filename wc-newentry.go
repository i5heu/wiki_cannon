package main

import (
	"fmt"
	"net/http"
	"time"
)

func NewentryHandler(w http.ResponseWriter, r *http.Request) {
	guestmodechek(w, r)

	t := "login: false"

	if checkLogin(r) == true {
		//if true == true {
		t = "login: true"

		newPublic := "1"
		newTitle := r.FormValue("Title")
		newNamepace := r.FormValue("Namespace")
		newPublic = r.FormValue("Public")
		newText := r.FormValue("Text")

		if len(newPublic) == 0 {

			newPublic = "1"
		}

		if newNamepace == "" {
			newNamepace = "main"
		}
		if newTitle == "" {
			t := time.Now()
			newTitle = t.String()[:len(t.String())-21]
		}

		db.Exec("INSERT INTO article(needlogin,title,namespace,text) VALUES(?,?,?,?)", newPublic, ReplaceSpecialChars(newTitle), ReplaceSpecialChars(newNamepace), newText)
		eventname := "ADD >" + ReplaceSpecialChars(newNamepace) + "/" + ReplaceSpecialChars(newTitle) + "< to articles"
		Eventloger(eventname, "wc-newentry", 0)
		checkErr(err)

		http.ServeFile(w, r, "./template/newentry.html")

	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}
