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
		newTags := r.FormValue("Tags")

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

		ids, err := db.Query("SELECT title , namespace FROM article WHERE title = ? AND namespace = ?", ReplaceSpecialChars(newTitle), ReplaceSpecialChars(newNamepace))
		checkErr(err)
		defer ids.Close()
		var namespace string
		var title string
		ids.Next()
		_ = ids.Scan(&title, &namespace)

		exist := false

		if (len(title) + len(namespace)) > 0 {
			exist = true
			t := time.Now()
			newTitle = newTitle + "-" + t.Format("2006-01-02-15-04-05")
		}

		db.Exec("INSERT INTO article(needlogin,title,namespace,text,tags,editcounter) VALUES(?,?,?,?,?,?)", newPublic, ReplaceSpecialChars(newTitle), ReplaceSpecialChars(newNamepace), newText, ReplaceSpecialChars(newTags), "1")
		eventname := "ADD >" + ReplaceSpecialChars(newNamepace) + "/" + ReplaceSpecialChars(newTitle) + "< to articles"
		Eventloger(eventname, "wc-newentry", 0)
		checkErr(err)

		if exist == false {
			http.ServeFile(w, r, "./template/newentry.html")
		} else {
			fmt.Fprintf(w, `Article already exist - title was chaged to curent Timestamp`)
		}
		refreshCache()
	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}
