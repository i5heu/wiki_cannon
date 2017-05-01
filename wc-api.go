package main

import (
	"fmt"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	t := "login: false"

	if checkLogin(r) == true {
		//if true == true {
		t = "login: true"

		newID := r.FormValue("Id")
		newNamepace := r.FormValue("Namespace")
		newTitle := r.FormValue("Title")
		newText := r.FormValue("Text")
		db.Exec("UPDATE `article` SET `namespace` = ?, `title` = ?, `text` = ? WHERE `article`.`id` = ? ", ReplaceSpecialChars(newNamepace), ReplaceSpecialChars(newTitle), newText, newID)

		checkErr(err)

		http.ServeFile(w, r, "newentry.html")

	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}
