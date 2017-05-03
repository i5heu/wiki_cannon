package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type view struct {
	Articlename string
	Path        string
	Title       template.HTML
	Text        template.HTML
}

var templatesView = template.Must(template.ParseFiles("./template/view.html", HtmlStructHeader, HtmlStructFooter))

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.Path)

	checkErr(err)
	encodetpath1 := strings.Split(u.Path, "/")

	if len(encodetpath1) < 4 {
		fmt.Fprintf(w, "ERROR 404")
		return
	}

	ids, err := db.Query("SELECT id,needlogin,namespace,title,text FROM article WHERE title=(?) AND namespace=(?)", encodetpath1[3], encodetpath1[2])
	checkErr(err)

	ids.Next()
	var id int
	var needlogin bool
	var namespace string
	var title string
	var text string
	_ = ids.Scan(&id, &needlogin, &namespace, &title, &text)
	checkErr(err)

	if needlogin == true && checkLogin(r) == false {
		fmt.Fprintf(w, "ERROR YOU ARE NOT LOGED IN")
		return
	}

	if id == 0 {
		fmt.Fprintf(w, "ERROR 404")
	} else {

		title = namespace + "/" + title

		TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title)))
		TextTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(text))))

		views := view{encodetpath1[2], title, TitleTMP, TextTMP}
		templatesView.Execute(w, views)
	}

}
