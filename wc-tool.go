package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
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

	if cookie == "PASSWORD" {
		return true
	}
	return false

}

////////////////////////////
func ReplaceSpecialChars(s string) (sc string) {
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "<", ">", "/", "#", "?", "=", "ß", "*", "'", "´", "\"", "%", ";", ":", "&", " "}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	sc = re.ReplaceAllString(s, "-")
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

	if cookie == "GUESTPWD" {
		return
	}
	http.Redirect(w, r, "/static/guestlogin.html", 302)
	return

}

var timemap = make(map[string]int64)
var TMPCACHECACHE = make(map[string]template.HTML)

func racepreventer(foo string) (bar template.HTML) {
	TMPCACHEWRITE = true
	/*if TMPCACHEWRITE == false {
	bar = TMPCACHE[foo]
	*/

	if TMPCACHECACHEWRITE == false {
		bar = TMPCACHECACHE[foo]
	}
	TMPCACHEWRITE = false
	return
}

var timeracer int8 = 0
var timenowracer int8 = 0

func timer(foo string) (a bool) {
	if timecachewrite == false {
		timecachewrite = true
		if timemap[foo] == 0 {
			timemap[foo] = int64(time.Now().Unix())
			timecachewrite = false
			return true
			fmt.Println("0")
		}
		timecachewrite = false
	}
	if timecachewrite == false {
		if timenowracer == 0 {
			timenowracer++
			if int64(time.Now().Unix()) > timemap[foo]+5 {
				timecachewrite = true
				timenowracer++
				timeracer++
				if timeracer == 0 {
					timeracer++
					timemap[foo] = int64(time.Now().Unix())
					timeracer = 0
					timenowracer = 0
				}
				timecachewrite = false
				return true
			}
			timenowracer = 0
		}

		return false
	} else {
		return false
	}
}
