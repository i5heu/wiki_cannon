package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	guestmodechek(w, r)
	u, err := url.Parse(r.URL.Path)
	checkErr(err)
	encodetpath1 := strings.Split(u.Path, "/")

	if len(encodetpath1) < 3 {
		fmt.Fprintf(w, "NO")
		return
	}

	switch encodetpath1[2] {
	case "":
		fmt.Fprintf(w, "NO METHOD SELECTED")
	case "editpost":
		ArticleEdit(w, r)
	default:
		fmt.Fprintf(w, "NO WORKING MONKEYS")
	}

}

func ArticleEdit(w http.ResponseWriter, r *http.Request) {
	t := "login: false"

	if checkLogin(r) == true {
		//if true == true {
		t = "login: true"

		newPublic := "1"
		newID := r.FormValue("Id")
		newNamepace := r.FormValue("Namespace")
		newPublic = r.FormValue("Public")
		newTitle := r.FormValue("Title")
		newText := r.FormValue("Text")

		if len(newPublic) == 0 {
			newPublic = "1"
		}

		db.Exec("INSERT INTO BUarticle(id,timec,timelastedit,needlogin,namespace,title,text,tags,viewcounter,editcounter)SELECT id,timec,timelastedit,needlogin,namespace,title,text,tags,viewcounter,editcounter FROM article WHERE id = ?", newID)

		db.Exec("UPDATE `article` SET `namespace` = ?, `title` = ?, `text` = ?, `needlogin` = ? WHERE `article`.`id` = ? ", ReplaceSpecialChars(newNamepace), ReplaceSpecialChars(newTitle), newText, newPublic, newID)

		eventname := "UPDATE >" + ReplaceSpecialChars(newNamepace) + "/" + ReplaceSpecialChars(newTitle) + "< to articles"
		eventID, _ := strconv.Atoi(newID)
		Eventloger(eventname, "wc-newentry", eventID)

		checkErr(err)

		http.ServeFile(w, r, "./template/newentry.html")

	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}
