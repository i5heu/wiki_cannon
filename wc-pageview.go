package main

import (
	"database/sql"
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

type NamespaceResult struct {
	ArticleId    int
	Articletitle template.HTML
	ArticleText  template.HTML
}

type namespace struct {
	Searchterm       string
	NamespaceResults []NamespaceResult
}

var namespaceView = template.Must(template.ParseFiles("./template/namespace.html", HtmlStructHeader, HtmlStructFooter))
var templatesView = template.Must(template.ParseFiles("./template/view.html", HtmlStructHeader, HtmlStructFooter))

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	guestmodechek(w, r)
	peageview++

	u, err := url.Parse(r.URL.Path)

	checkErr(err)
	encodetpath1 := strings.Split(u.Path, "/")

	if len(encodetpath1) == 3 {
		NamespaceHandler(w, r)
		return
	}
	if len(encodetpath1) < 4 {
		fmt.Fprintf(w, "ERROR 404")
		return
	}

	ids, err := db.Query("SELECT id,needlogin,namespace,title,text FROM article WHERE title=(?) AND namespace=(?)", encodetpath1[3], encodetpath1[2])
	defer ids.Close()
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

var tmpNamespace []NamespaceResult

func NamespaceHandler(w http.ResponseWriter, r *http.Request) {

	tmpNamespace = tmpNamespace[:0]

	searchterm := r.URL.Path[3:]
	if len(searchterm) == 0 {
		http.Redirect(w, r, "/desk", 302)
		return
	}

	newquery := searchterm

	var ids *sql.Rows

	switch searchterm {
	case "":
		http.Redirect(w, r, "/desk", 302)

	default:
		ids, err = db.Query("SELECT  id,namespace,title,SUBSTR(text,1,100) FROM article WHERE (needlogin = '0' OR needlogin = ?) AND namespace = ? ORDER BY timec DESC LIMIT 100", checkLogin(r), newquery)
		defer ids.Close()
	}

	checkErr(err)
	for ids.Next() {
		var id int
		var namespace string
		var title string
		var text string
		_ = ids.Scan(&id, &namespace, &title, &text)
		checkErr(err)
		title = namespace + "/" + title
		text = text + "..."

		TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title)))
		TextTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(text))))
		tmpNamespace = append(tmpNamespace, NamespaceResult{id, TitleTMP, TextTMP})
	}
	searchs := namespace{searchterm, tmpNamespace}

	namespaceView.Execute(w, searchs)
}
