package main

import (
	"html/template"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Article struct {
	ArticleId        int
	Url              template.HTML
	ArticleNamespace template.HTML
	Articletitle     template.HTML
	ArticleText      template.HTML
}

type lista struct {
	Login     bool
	LoginText string
	Articles  []Article
	Geldlog   template.HTML
}

var templatesDesktop = template.Must(template.ParseFiles("./template/desktop.html", HtmlStructHeader, HtmlStructFooter))
var GeldlogTMPCACHE template.HTML
var tmp []Article

func DesktopHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler
	guestmodechek(w, r)

	login := false

	//if int64(time.Now().Unix()) > timecache+5 {
	cache(checkLogin(r))
	//}

	t := "login: false"
	if checkLogin(r) == true {
		t = "login: true"
		login = true
	}

	lists := lista{login, t, tmp, Geldlogfunc()}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		templatesDesktop.Execute(w, lists)
	}
}

func cache(login bool) {
	tmp = tmp[:0]
	ids, err := db.Query("SELECT id, namespace, title, LEFT (text,100) FROM `article` WHERE (needlogin = '0' OR needlogin = ?) ORDER BY id DESC LIMIT 10", login)
	checkErr(err)

	for ids.Next() {
		var id int
		var namespace string
		var title string
		var text string
		_ = ids.Scan(&id, &namespace, &title, &text)
		checkErr(err)

		text = text + "..."
		url := namespace + "/" + title

		UrlTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(url)))
		NamespaceTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(namespace)))
		TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title)))
		TextTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(text))))

		tmp = append(tmp, Article{id, UrlTMP, NamespaceTMP, TitleTMP, TextTMP})

	}

}

func Geldlogfunc() (GeldlogTMP template.HTML) {
	ids, err := db.Query("SELECT title FROM `article`")
	checkErr(err)

	GeldlogTMP = template.HTML("<h1>TEST</h1>")

	for ids.Next() {
		var title string
		_ = ids.Scan(&title)
		checkErr(err)

		GeldlogTMP += template.HTML("<b>") + template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title))) + template.HTML("</b> <br>")
		GeldlogTMPCACHE = GeldlogTMP
	}

	return
}
