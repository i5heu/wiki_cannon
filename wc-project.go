package main

import (
	"html/template"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

type ProjectResult struct {
	ProjectId     int
	ProjectTitle1 template.HTML
	ProjectTitle2 template.HTML
	ProjectTags1  template.HTML
	ProjectNum1   int
	ProjectNum2   int
}

type pro struct {
	Test           template.HTML
	ProjectResults []ProjectResult
}

var templatesProject = template.Must(template.ParseFiles("./template/project.html", HtmlStructHeader, HtmlStructFooter))

func ProjectHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler
	guestmodechek(w, r)

	ids, err := db.Query("SELECT  ItemID,title1,title2,tags1,num1,num2 FROM items WHERE APP='project' ORDER BY timecreate DESC LIMIT 100")
	defer ids.Close()
	checkErr(err)
	var ProjectTMP []ProjectResult

	for ids.Next() {
		var id int
		var title1 string
		var title2 string
		var tags1 string
		var num1 int
		var num2 int
		_ = ids.Scan(&id, &title1, &title2, &tags1, &num1, &num2)
		checkErr(err)

		Title1TMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title1)))
		Title2TMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(title2)))
		Tags1TMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte(tags1)))

		ProjectTMP = append(ProjectTMP, ProjectResult{id, Title1TMP, Title2TMP, Tags1TMP, num1, num2})
	}

	lists := pro{template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte("HELLO"))), ProjectTMP}

	templatesProject.Execute(w, lists)

}
