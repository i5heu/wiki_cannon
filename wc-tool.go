package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

///////////////////////////
func sessionExists(r *http.Request, cookiename string) bool {
	_, err := r.Cookie(cookiename)
	if err == http.ErrNoCookie {
		return false
	} else if err != nil {
		log.Println(err)
		return false
	}

	return true
}

/////////////////////
func checkLogin(r *http.Request) bool {
	var cookie string
	var cookieTMP *http.Cookie

	if sessionExists(r, "pwd") == true {
		cookieTMP, _ = r.Cookie("pwd")
		cookie = cookieTMP.Value
	} else {
		return false
	}

	if cookie == personalpwd {
		return true
	}
	return false

}

////////////////////////////
func ReplaceSpecialChars(s string) (sc string) {
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "<", ">", "+", "/", "#", "?", "=", "ß", "*", "'", "´", "\"", "%", ";", ":", "&", " "}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	sc = re.ReplaceAllString(s, "-")
	return
}

func ReplaceSpecialCharsWith_(s string) (sc string) {
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "<", ">", "+", "-", "/", "#", "?", "=", "ß", "*", "'", "´", "\"", "%", ";", ":", "&", " "}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	sc = re.ReplaceAllString(s, "_")
	return
}

func ReplaceSpecialCharsWithSpaceSpaceALLOWED(s string) (sc string) {
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "<", ">", "+", "-", "/", "#", "?", "=", "ß", "*", "'", "´", "\"", "%", ";", ":", "&", "\n"}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	sc = re.ReplaceAllString(s, " ")
	return
}

func Eventloger(name string, changeAPP string, changeID int) {
	db.Exec("INSERT INTO eventlog(name,changeAPP,changeID) VALUES(?,?,?)", name, changeAPP, changeID)
}

///////////////////
func guestmodechek(w http.ResponseWriter, r *http.Request) {

	if guestmode == false {
		return
	}

	var cookie string
	var cookieTMP *http.Cookie

	if sessionExists(r, "pwdguest") == true {
		cookieTMP, _ = r.Cookie("pwdguest")
		cookie = cookieTMP.Value
	} else {
		http.Redirect(w, r, "/static/guestlogin.html", 302)
		return
	}

	if cookie == guestpwd {
		return
	}
	http.Redirect(w, r, "/static/guestlogin.html", 302)
	return

}

func numberswithcoma(foo string) (bar string) {
	if len(foo) < 2 {
		bar = foo
		return
	} else {
		if len(foo) < 3 {
			foo = "0" + foo
		}
		bar = foo[:len(foo)-2] + "," + foo[len(foo)-2:]
		return
	}
}

func CheckIfOnlyNumbers(w http.ResponseWriter, r *http.Request, v string) (re bool) {
	if _, err := strconv.Atoi(v); err == nil {
		re = true
		return
	} else {
		re = false
		fmt.Fprintf(w, "%q looks like not a number.\n", v)
		return
	}
}

func ItemBackuper(id int) {
	db.Exec("INSERT INTO BUitems(ItemID,timecreate,timelastedit,needlogin,APP,viewcounter,editcounter,title1,title2,text1,text2,tags1,num1,num2,num3)SELECT ItemID,timecreate,timelastedit,needlogin,APP,viewcounter,editcounter,title1,title2,text1,text2,tags1,num1,num2,num3 FROM items WHERE ItemID = ?", id)
}
