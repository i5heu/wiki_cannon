package main

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type SearchResult struct {
	ArticleId    int
	Articletitle template.HTML
	ArticleText  template.HTML
}

type search struct {
	Searchterm   string
	SeachResults []SearchResult
}

var templatesSearch = template.Must(template.ParseFiles("search.html", HtmlStructFooter, HtmlStructHeader))
var tmpSearch []SearchResult

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tmpSearch = tmpSearch[:0]
	searchterm := r.URL.Path[3:]

	newquery := "*" + searchterm + "*"

	var ids *sql.Rows

	switch searchterm {
	case "all":
		ids, err = db.Query("SELECT  id,namespace,title,SUBSTR(text,1,100) FROM article ORDER BY timec DESC")
	default:
		ids, err = db.Query("SELECT  id,namespace,title,SUBSTR(text,1,100) FROM article WHERE MATCH (title,text) AGAINST (? IN BOOLEAN MODE)", newquery)
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
		tmpSearch = append(tmpSearch, SearchResult{id, TitleTMP, TextTMP})
	}
	searchs := search{searchterm, tmpSearch}

	templatesSearch.Execute(w, searchs)

}
