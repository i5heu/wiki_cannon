package main

import (
	"fmt"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	guestmodechek(w, r)
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

		db.Exec("UPDATE `article` SET `namespace` = ?, `title` = ?, `text` = ? WHERE `article`.`id` = ?, `needlogin` = ? ", ReplaceSpecialChars(newNamepace), ReplaceSpecialChars(newTitle), newText, newID, newPublic)

		checkErr(err)

		http.ServeFile(w, r, "./template/newentry.html")

	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}
