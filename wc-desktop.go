package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/microcosm-cc/bluemonday"
)

type lista struct {
	Login     bool
	LoginText string
	Articles  template.HTML
	Geldlog   template.HTML
}

var templatesDesktop = template.Must(template.ParseFiles("./template/desktop.html", HtmlStructHeader, HtmlStructFooter))
var TMPCACHE = make(map[string]template.HTML)
var TMPCACHECACHE = make(map[string]template.HTML)
var TMPCACHEWRITE bool = false
var TMPCACHECACHEWRITE bool = false
var peageview int64

func DesktopHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler
	guestmodechek(w, r)

	login := false
	cachetimername := "article-" + strconv.FormatBool(checkLogin(r))
	cachegeldlogname := "geldlog-" + strconv.FormatBool(checkLogin(r))

	t := "login: false"
	if checkLogin(r) == true {
		t = "login: true"
		login = true
	}

	lists := lista{}

	if TMPCACHEWRITE == false {
		lists = lista{login, t, TMPCACHE[cachetimername], TMPCACHE[cachegeldlogname]}
	} else {
		if TMPCACHECACHEWRITE == true {
			lists = lista{login, t, TMPCACHECACHE[cachetimername], TMPCACHECACHE[cachegeldlogname]}
		}
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		templatesDesktop.Execute(w, lists)
		peageview++
	}
}

func cache(login bool, foo string) {
	TMPCACHE[foo] = template.HTML("<h1>TEST</h1><br><hr>")
	ids, err := db.Query("SELECT id, namespace, title FROM `article` WHERE (needlogin = '0' OR needlogin = ?) ORDER BY id DESC LIMIT 10", login)

	checkErr(err)

	for ids.Next() {
		var id string
		var namespace string
		var title string

		_ = ids.Scan(&id, &namespace, &title)
		checkErr(err)
		url := namespace + "/" + title

		UrlTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(url)))
		NamespaceTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(namespace)))
		TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title)))
		TMPCACHE[foo] += template.HTML("<b>") + template.HTML(id) + template.HTML("</b>  ") + template.HTML("<a href='/p/") + UrlTMP + template.HTML("'>") + NamespaceTMP + template.HTML("/") + TitleTMP + template.HTML("</a><br>\n")

	}
}

func Geldlogfunc(foo string) (GeldlogTMP template.HTML) {

	ids, err := db.Query("SELECT title FROM `article` ORDER by ID DESC LIMIT 15")
	checkErr(err)

	GeldlogTMP = template.HTML("<h1>TEST</h1>")

	for ids.Next() {
		var title string
		_ = ids.Scan(&title)
		checkErr(err)

		GeldlogTMP += template.HTML("<b>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title))) + template.HTML("</b> <br>")

	}

	TMPCACHE[foo] = GeldlogTMP
	return
}
