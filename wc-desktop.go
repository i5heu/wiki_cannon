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
	Eventlog  template.HTML
}

var templatesDesktop = template.Must(template.ParseFiles("./template/desktop.html", HtmlStructHeader, HtmlStructFooter))
var TMPCACHE = make(map[string]template.HTML)
var TMPCACHECACHE = make(map[string]template.HTML)
var TMPCACHEWRITE bool = false
var TMPCACHECACHEWRITE bool = false

func DesktopHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler
	guestmodechek(w, r)

	login := false
	cachetimername := "article-" + strconv.FormatBool(checkLogin(r))
	cachegeldlogname := "geldlog-" + strconv.FormatBool(checkLogin(r))
	cacheeventname := "event-" + strconv.FormatBool(checkLogin(r))

	t := "login: false"
	if checkLogin(r) == true {
		t = "login: true "
		login = true
	}
	lists := lista{}

	if TMPCACHEWRITE == false {
		lists = lista{login, t, TMPCACHE[cachetimername], TMPCACHE[cachegeldlogname], TMPCACHECACHE[cacheeventname]}
	} else if TMPCACHECACHEWRITE == true {
		lists = lista{login, t, TMPCACHECACHE[cachetimername], TMPCACHECACHE[cachegeldlogname], TMPCACHECACHE[cacheeventname]}
	} else {
		lists = lista{login, "PLEASE RELOAD", template.HTML("<b>Please reload this page</b>"), template.HTML("<b>Please reload this page</b>"), template.HTML("<b>Please reload this page</b>")}
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		templatesDesktop.Execute(w, lists)
	}
}

func cache(login bool, foo string) {
	TMPCACHE[foo] = template.HTML("Last Article<br>----------<br>")
	ids, err := db.Query("SELECT id, namespace, title FROM `article` WHERE (needlogin = '0' OR needlogin = ?) ORDER BY id DESC LIMIT 15", login)
	defer ids.Close()
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
		TMPCACHE[foo] += template.HTML(`<tr><td class='borderfull'>`) + template.HTML(id) + template.HTML("</td><td class='borderfull'>") + template.HTML("<a href='/p/") + UrlTMP + template.HTML("'>") + NamespaceTMP + template.HTML("/") + TitleTMP + template.HTML("</a></td></tr>\n")

	}

}

func Geldlogfunc(foo string) (GeldlogTMP template.HTML) {

	ids, err := db.Query("SELECT SUM(num1) FROM `items` WHERE APP='geldlog' AND timecreate >= ( CURDATE() - INTERVAL 30 DAY )")
	defer ids.Close()
	ids.Next()
	var sume string
	_ = ids.Scan(&sume)
	sume = numberswithcoma(sume)
	GeldlogTMP = template.HTML("Geldlog<br>----------<br>") + template.HTML(sume) + template.HTML("€ sum of last 30Days")

	ids, err = db.Query("SELECT title1, num1, DATEDIFF(CURDATE(),timecreate) FROM `items` WHERE APP='geldlog' AND timecreate >= ( CURDATE() - INTERVAL 3 DAY ) ORDER by timecreate DESC LIMIT 13")
	checkErr(err)

	for ids.Next() {
		var title1 string
		var num1 string
		var daysago string
		_ = ids.Scan(&title1, &num1, &daysago)
		checkErr(err)
		num1 = numberswithcoma(num1)
		GeldlogTMP += template.HTML("<tr><td class='borderfull'>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(daysago))) + template.HTML("</td><td class='borderfull'>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title1))) + template.HTML("</td><td class='borderfull'> ") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(num1))) + template.HTML("</td></tr>")
	}

	TMPCACHE[foo] = GeldlogTMP
	return
}

func Eventlogfunc(foo string) (EventlogTMP template.HTML) {

	ids, err := db.Query("SELECT id,name,changeAPP num1 FROM `eventlog` ORDER by time DESC LIMIT 20")
	defer ids.Close()
	checkErr(err)

	for ids.Next() {
		var id string
		var name string
		var changeAPP string
		_ = ids.Scan(&id, &name, &changeAPP)
		checkErr(err)

		EventlogTMP += template.HTML("<tr><td class='borderfull'>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(id))) + template.HTML("</td><td class='borderfull'>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(name))) + template.HTML("</td><td class='borderfull'> ") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(changeAPP))) + template.HTML("</td></tr>")
	}

	TMPCACHE[foo] = EventlogTMP
	return
}
