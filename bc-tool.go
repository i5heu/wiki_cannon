package main

import (
	"log"
	"net/http"
	"regexp"
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
