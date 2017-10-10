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
	ArticleID    int
	Articlename  string
	Path         string
	Title        template.HTML
	Tags         string
	Text         template.HTML
	Viewcounter  int
	Editcounter  int
	DarkTemplate bool
}

type NamespaceResult struct {
	ArticleId    int
	Articletitle template.HTML
	ArticleText  template.HTML
}

type namespace struct {
	Searchterm       string
	NamespaceResults []NamespaceResult
	DarkTemplate     bool
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	guestmodechek(w, r)

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

	ids, err := db.Query("SELECT id,needlogin,namespace,title,tags,text,viewcounter,editcounter FROM article WHERE title=(?) AND namespace=(?)", encodetpath1[3], encodetpath1[2])
	defer ids.Close()
	checkErr(err)

	ids.Next()
	var id int
	var needlogin bool
	var namespace string
	var title string
	var tags string
	var text string
	var viewcounter int
	var editcounter int
	_ = ids.Scan(&id, &needlogin, &namespace, &title, &tags, &text, &viewcounter, &editcounter)
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

		DarkTemplate := ChekDarkTemplate(r)

		views := view{id, encodetpath1[2], title, TitleTMP, tags, TextTMP, viewcounter, editcounter, DarkTemplate}
		templatesView.Execute(w, views)

		db.Exec("UPDATE `article` SET viewcounter = IFNULL(`viewcounter`, 0) + 1 WHERE id = ?", id)

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

	DarkTemplate := ChekDarkTemplate(r)

	searchs := namespace{searchterm, tmpNamespace, DarkTemplate}

	namespaceView.Execute(w, searchs)
}
