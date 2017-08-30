package main

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

type SearchResult struct {
	ArticleId    int
	Articletitle template.HTML
	ArticleTags  template.HTML
}

type search struct {
	Searchterm   string
	SeachResults []SearchResult
}

var tmpSearch []SearchResult

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	guestmodechek(w, r)

	tmpSearch = tmpSearch[:0]

	searchterm := r.URL.Path[3:]

	if len(searchterm) == 0 {
		http.Redirect(w, r, "/desk", 302)
		return
	}

	searchterm = ReplaceSpecialCharsWith_(searchterm)

	newquery := "%" + searchterm + "%"

	var ids *sql.Rows

	switch searchterm {
	case "all":
		ids, err = db.Query("SELECT  id,namespace,title,tags FROM article ORDER BY timec DESC ORDER BY timelastedit DESC LIMIT 200")
		defer ids.Close()
	case "":
		http.Redirect(w, r, "/desk", 302)

	default:
		ids, err = db.Query("SELECT  id,namespace,title,SUBSTR(tags,1,100) FROM article WHERE (needlogin = '0' OR needlogin = ?) AND CONCAT(title,tags,namespace) LIKE ? ORDER BY timelastedit DESC LIMIT 200", checkLogin(r), newquery)
		defer ids.Close()
	}

	checkErr(err)
	for ids.Next() {
		var id int
		var namespace string
		var title string
		var tags string
		_ = ids.Scan(&id, &namespace, &title, &tags)
		checkErr(err)
		title = namespace + "/" + title
		tags = tags + "…"

		TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title)))
		TagsTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(tags)))
		tmpSearch = append(tmpSearch, SearchResult{id, TitleTMP, TagsTMP})
	}
	searchs := search{searchterm, tmpSearch}

	templatesSearch.Execute(w, searchs)
}
