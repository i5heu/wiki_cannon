package main

import (
	"fmt"
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
var TMPCACHEREAD bool = false
var TMPCACHEWRITE bool = false
var TMPCACHECACHEWRITE bool = false
var timecachewrite bool = false
var reacewriteer int8 = 0

func DesktopHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler

	guestmodechek(w, r)

	login := false
	cachetimername := "article-" + strconv.FormatBool(checkLogin(r))
	cachegeldlogname := "geldlog-" + strconv.FormatBool(checkLogin(r))

	if TMPCACHEWRITE == false {
		if timecachewrite == false {
			if timer(cachetimername) == true {
				if TMPCACHEREAD == false {
					if TMPCACHECACHEWRITE == false {
						TMPCACHECACHEWRITE = true
						reacewriteer++
						cache(checkLogin(r), cachetimername)
						TMPCACHECACHEWRITE = false
					}
				}
			}
		}
		if timecachewrite == false {
			if TMPCACHEREAD == false {
				if timer(cachegeldlogname) == true {
					if TMPCACHECACHEWRITE == false {
						TMPCACHECACHEWRITE = true
						Geldlogfunc(cachegeldlogname)

						TMPCACHECACHEWRITE = false
					}
				}
			}
		}
		TMPCACHEWRITE = true
		TMPCACHE = TMPCACHECACHE
		TMPCACHEWRITE = false
	}

	t := "login: false"
	if checkLogin(r) == true {
		t = "login: true"
		login = true
	}

	lists := lista{login, t, racepreventer(cachetimername), racepreventer(cachegeldlogname)}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		templatesDesktop.Execute(w, lists)
	}
}

func cache(login bool, foo string) {
	fmt.Println(reacewriteer)
	if reacewriteer == 1 {
		reacewriteer++
		fmt.Println(">", reacewriteer)
		TMPCACHECACHE[foo] = template.HTML("<h1>TEST</h1><br><hr>")
	}
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
		TMPCACHECACHE[foo] += template.HTML("<b>") + template.HTML(id) + template.HTML("</b>  ") + template.HTML("<a href='/p/") + UrlTMP + template.HTML("'>") + NamespaceTMP + template.HTML("/") + TitleTMP + template.HTML("</a><br>\n")

	}
	reacewriteer = 0
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
	TMPCACHECACHE[foo] = GeldlogTMP
	return
}
