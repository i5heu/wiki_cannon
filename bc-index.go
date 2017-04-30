package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Article struct {
	ArticleId    int
	Articletitle template.HTML
	ArticleText  template.HTML
}

type lista struct {
	Login     bool
	LoginText string
	Articles  []Article
}

const (
	timeFormat = "2006-01-02 15:04 MST"
)

var templatesIndex = template.Must(template.ParseFiles("index.html", HtmlStructHeader, HtmlStructFooter))
var timecache int64 = time.Now().Unix() - 10
var tmp []Article

func IndexHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler
	login := false

	if int64(time.Now().Unix()) > timecache+5 {
		cache()
	}

	t := "login: false"
	if checkLogin(r) == true {
		t = "login: true"
		login = true
	}

	lists := lista{login, t, tmp}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		templatesIndex.Execute(w, lists)
	}
}

func cache() {
	tmp = tmp[:0]

	ids, err := db.Query("SELECT id, namespace, title, LEFT (text,200) FROM `article` ORDER BY id DESC LIMIT 5")
	checkErr(err)

	for ids.Next() {
		var id int
		var namespace string
		var title string
		var text string
		_ = ids.Scan(&id, &namespace, &title, &text)
		checkErr(err)

		text = text + "..."
		title = namespace + "/" + title

		TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title)))
		TextTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(text))))

		tmp = append(tmp, Article{id, TitleTMP, TextTMP})

	}

	timecache = time.Now().Unix()
}
