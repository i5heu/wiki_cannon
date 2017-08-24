package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

type lista struct {
	Login         bool
	LoginText     string
	CurentVersion string
	Update        bool
	UpdateVersion string
	Articles      template.HTML
	Geldlog       template.HTML
	Eventlog      template.HTML
	Project       template.HTML
	Namespace     template.HTML
	Shortcut      template.HTML
	Lastedit      template.HTML
	Rendertime    time.Duration
}

var templatesDesktop = template.Must(template.ParseFiles("./template/desktop.html", HtmlStructHeader, HtmlStructFooter))
var TMPCACHE = make(map[string]template.HTML)
var TMPCACHECACHE = make(map[string]template.HTML)
var TMPCACHEWRITE bool = false
var TMPCACHECACHEWRITE bool = false

func DesktopHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler
	start := time.Now()
	guestmodechek(w, r)

	login := false
	cachetimername := "article-" + strconv.FormatBool(checkLogin(r))
	cachegeldlogname := "geldlog-" + strconv.FormatBool(checkLogin(r))
	cacheeventname := "event-" + strconv.FormatBool(checkLogin(r))
	namespacename := "namespace-" + strconv.FormatBool(checkLogin(r))
	projectname := "project-" + strconv.FormatBool(checkLogin(r))
	lasteditname := "lastedit-" + strconv.FormatBool(checkLogin(r))
	shortcut := "shortcut-" + strconv.FormatBool(checkLogin(r))

	t := "login: false"
	if checkLogin(r) == true {
		t = "login: true "
		login = true
	}
	lists := lista{}

	if TMPCACHEWRITE == false {
		lists = lista{login, t, wcversion, WcVersionUpdateBOOL, WcVersionUpdate, TMPCACHE[cachetimername], TMPCACHE[cachegeldlogname], TMPCACHE[cacheeventname], TMPCACHE[projectname], TMPCACHE[namespacename], TMPCACHE[shortcut], TMPCACHE[lasteditname], time.Since(start)}
	} else if TMPCACHECACHEWRITE == false {
		lists = lista{login, t, wcversion, WcVersionUpdateBOOL, WcVersionUpdate, TMPCACHECACHE[cachetimername], TMPCACHECACHE[cachegeldlogname], TMPCACHECACHE[cacheeventname], TMPCACHE[projectname], TMPCACHECACHE[namespacename], TMPCACHECACHE[shortcut], TMPCACHECACHE[lasteditname], time.Since(start)}
	} else {
		lists = lista{login, "PLEASE RELOAD", wcversion, WcVersionUpdateBOOL, WcVersionUpdate, template.HTML("<b>Please reload this page</b>"), template.HTML("<b>Please reload this page</b>"), template.HTML("<b>Please reload this page</b>"), template.HTML("<b>Please reload this page</b>"), template.HTML("<b>Please reload this page</b>"), template.HTML("<b>Please reload this page</b>"), template.HTML("<b>Please reload this page</b>"), time.Since(start)}
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		templatesDesktop.Execute(w, lists)
	}
	fmt.Println("DESK:", time.Since(start))
}

func cache(login bool, foo string) {
	TMPCACHE[foo] = template.HTML("Last New Article<br>----------<br>")
	ids, err := db.Query("SELECT id, namespace, title FROM `article` WHERE (needlogin = '0' OR needlogin = ?) ORDER BY id DESC LIMIT 25", login)
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

func Geldlogfunc(foo string) {

	var GeldlogTMP template.HTML

	ids, err := db.Query("SELECT SUM(num1) FROM `items` WHERE APP='geldlog' AND MONTH(timecreate) = MONTH(NOW())")
	defer ids.Close()
	ids.Next()
	var sume string
	_ = ids.Scan(&sume)
	sume = numberswithcoma(sume)
	GeldlogTMP = template.HTML("Geldlog<br>----------<br>") + template.HTML(sume) + template.HTML("€ sum of curent Month")

	ids, err = db.Query("SELECT title1, num1, DATEDIFF(CURDATE(),timecreate) FROM `items` WHERE APP='geldlog' AND timecreate >= ( CURDATE() - INTERVAL 7 DAY ) ORDER by timecreate DESC LIMIT 50")
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

func Eventlogfunc(foo string) {
	var EventlogTMP template.HTML

	ids, err := db.Query("SELECT id,name,changeAPP num1 FROM `eventlog` ORDER by time DESC LIMIT 25")
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

func Namespacefunc(foo string) {
	var NamespaceTMP template.HTML

	ids, err := db.Query("SELECT DISTINCT(namespace) FROM article ORDER BY namespace ASC;")
	defer ids.Close()
	checkErr(err)

	for ids.Next() {
		var namespace string
		_ = ids.Scan(&namespace)
		checkErr(err)

		NamespaceTMP += template.HTML(`<tr><td class='borderfull'><a href="/p/`) + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(namespace))) + template.HTML(`">`) + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(namespace))) + template.HTML("</a></td></tr>")
	}

	TMPCACHE[foo] = NamespaceTMP
	return
}

func Lasteditfunc(foo string) {
	var TMP template.HTML

	ids, err := db.Query("SELECT title, namespace FROM article ORDER BY timelastedit DESC LIMIT 50;")
	defer ids.Close()
	checkErr(err)

	for ids.Next() {
		var title, namespace string
		_ = ids.Scan(&title, &namespace)
		checkErr(err)

		url := namespace + "/" + title
		UrlTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(url)))

		TMP += template.HTML(`<tr><td class='borderfull'><a href="/p/`) + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(UrlTMP))) + template.HTML(`">`) + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(UrlTMP))) + template.HTML("</a></td></tr>")
	}

	TMPCACHE[foo] = TMP
	return
}

func Projectfunc(foo string) {
	var ProjectTMP template.HTML

	ids, err := db.Query("SELECT title1,title2,num2 FROM items WHERE APP='project' ORDER BY num2 DESC LIMIT 50")
	defer ids.Close()
	checkErr(err)

	for ids.Next() {
		var title1 string
		var title2 string
		var num2 string

		_ = ids.Scan(&title1, &title2, &num2)
		checkErr(err)

		ProjectTMP += template.HTML("<tr><td class='borderfull'>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title1))) + template.HTML("</td><td class='borderfull'>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title2))) + template.HTML("</td><td class='borderfull'>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(num2))) + template.HTML("</td></tr>")
	}

	TMPCACHE[foo] = ProjectTMP
}

func Shortcutfunc(foo string) {
	var ProjectTMP template.HTML

	ids, err := db.Query("SELECT title1,text1 FROM items WHERE APP='shortcut' ORDER BY num1 DESC")
	defer ids.Close()
	checkErr(err)

	for ids.Next() {
		var title1 string
		var text1 string

		_ = ids.Scan(&title1, &text1)
		checkErr(err)

		ProjectTMP += template.HTML("<tr><td class='borderfull'> <a href='") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(text1))) + template.HTML("'>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title1))) + template.HTML("</a></td></tr>")
	}

	TMPCACHE[foo] = ProjectTMP
}
