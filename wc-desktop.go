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
}

/*const (
	timeFormat = "2006-01-02 15:04 MST"
)
*/
var templatesDesktop = template.Must(template.ParseFiles("./template/desktop.html", HtmlStructHeader, HtmlStructFooter))

//var timecache int64 = time.Now().Unix() - 10
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

	lists := lista{login, t, tmp}

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

	//	timecache = time.Now().Unix()
}
