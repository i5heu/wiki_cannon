package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
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
	case "geldlog":
		AddGeldlog(w, r)
	case "authenticator":
		Authenticator(w, r)
	case "project":
		Project(w, r)
	case "projectDel":
		ProjectDel(w, r, encodetpath1[3])
	default:
		fmt.Fprintf(w, "NO WORKING MONKEYS")
		return
	}

}

func AddGeldlog(w http.ResponseWriter, r *http.Request) {
	if checkLogin(r) == false {
		fmt.Fprintf(w, `You have to login to do this!`)
		return
	}
	newTitle := r.FormValue("Title")
	newTitle2 := r.FormValue("Title2")
	newText := r.FormValue("Text")
	newTags := r.FormValue("Tags")
	newNum := r.FormValue("Num")

	if strings.ContainsAny(newNum, ",") == true {
		if strings.Index(newNum, ",") == len(newNum)-3 {
			newNum = strings.Replace(newNum, ",", "", -1)
		} else if strings.Index(newNum, ",") == len(newNum)-2 {
			newNum = strings.Replace(newNum, ",", "", -1)
			foo, _ := strconv.Atoi(newNum)
			newNum = strconv.Itoa(foo * 10)
		} else {
			fmt.Fprintf(w, `only 2 digits after "," alowed`)
			return
		}
	} else {
		foo, _ := strconv.Atoi(newNum)
		newNum = strconv.Itoa(foo * 100)
	}

	if newNum == "" {
		newNum = "0"
	}
	db.Exec("INSERT INTO `items` ( `APP`, `title1`, `title2`, `text1`, `tags1`, `num1`) VALUES (?,?,?,?,?,?);", "geldlog", newTitle, newTitle2, newText, newTags, newNum)

	eventname := "ADD >" + ReplaceSpecialChars(newTitle) + " - " + ReplaceSpecialChars(numberswithcoma(newNum)) + "€"
	Eventloger(eventname, "geldlog", 0)
	refreshCache()
	fmt.Fprintf(w, `<h1>WAS SEND!</h1> <meta http-equiv="refresh" content="0; url=/">	`)

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

		db.Exec("UPDATE `article` SET `timelastedit` = NOW() ,`namespace` = ?, `title` = ?, `text` = ?, `needlogin` = ?, `editcounter` = IFNULL(`editcounter`, 0) + 1  WHERE `id` = ? ", ReplaceSpecialChars(newNamepace), ReplaceSpecialChars(newTitle), newText, newPublic, newID)

		eventname := "UPDATE >" + ReplaceSpecialChars(newNamepace) + "/" + ReplaceSpecialChars(newTitle) + "< to articles"
		eventID, _ := strconv.Atoi(newID)
		Eventloger(eventname, "wc-newentry", eventID)

		checkErr(err)
		refreshCache()
		http.ServeFile(w, r, "./template/newentry.html")

	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}

func Authenticator(w http.ResponseWriter, r *http.Request) {

	UserInputTMP := r.FormValue("pwd")

	foo := sha256.Sum256([]byte(UserInputTMP))

	bar := hex.EncodeToString(foo[:])

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "pwd", Value: bar, Path: "/", Expires: expiration}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}

func Project(w http.ResponseWriter, r *http.Request) {
	if checkLogin(r) == false {
		fmt.Fprintf(w, `You have to login to do this!`)
		return
	}

	newTitle1 := r.FormValue("Title1")
	newTitle2 := r.FormValue("Title2")
	newTags1 := r.FormValue("Tags1")
	newNum1 := r.FormValue("Num1")
	newNum2 := r.FormValue("Num2")

	if CheckIfOnlyNumbers(w, r, newNum1) == false {
		return
	}
	if CheckIfOnlyNumbers(w, r, newNum2) == false {
		return
	}

	db.Exec("INSERT INTO `items` ( `APP`, `title1`, `title2`, `tags1`, `num1`, `num2`) VALUES (?,?,?,?,?,?);", "project", ReplaceSpecialChars(newTitle1), ReplaceSpecialChars(newTitle2), ReplaceSpecialChars(newTags1), newNum1, newNum2)

	refreshCache()

	http.Redirect(w, r, "/project/", 302)
	eventname := "ADD >" + ReplaceSpecialChars(newTitle1) + "< to Project"
	Eventloger(eventname, "project", 0)
}

func ProjectDel(w http.ResponseWriter, r *http.Request, idTMP string) {
	if checkLogin(r) == false {
		fmt.Fprintf(w, `You have to login to do this!`)
		return
	}

	id, err := strconv.Atoi(idTMP)
	if err != nil {
		fmt.Fprintf(w, `The Attribute is not convertible to an Int`)
	}

	ItemBackuper(id)
	eventname := "DEL >" + strconv.Itoa(id) + "< from Project"
	Eventloger(eventname, "project", id)

	db.Exec("DELETE from items WHERE ItemID = ?", id)

	refreshCache()

	http.Redirect(w, r, "/project/", 302)
	return
}
