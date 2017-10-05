package main

import (
	"html/template"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
)

type ProjectResult struct {
	ProjectId       int
	ProjectTitle1   template.HTML
	ProjectTitle2   template.HTML
	ProjectTags1    template.HTML
	ProjectNum1     int
	ProjectNum2     int
	ProjectRowClass string
}

type pro struct {
	Test           template.HTML
	ProjectResults []ProjectResult
	DarkTemplate   bool
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler
	guestmodechek(w, r)
	if checkLogin(r) == false {
		http.Redirect(w, r, "/desk", 302)
		return
	}

	ids, err := db.Query("SELECT  ItemID,title1,title2,tags1,num1,num2 FROM items WHERE APP='project' ORDER BY num2 DESC LIMIT 100")
	defer ids.Close()
	checkErr(err)
	var ProjectTMP []ProjectResult
	var projectClassSwitch bool = true
	var ProjectClassTMP = "ProjectTableDark"

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

		if projectClassSwitch == true {
			ProjectClassTMP = "ProjectTableDark"
			projectClassSwitch = false
		} else {
			ProjectClassTMP = "ProjectTableBright"
			projectClassSwitch = true
		}

		ProjectTMP = append(ProjectTMP, ProjectResult{id, Title1TMP, Title2TMP, Tags1TMP, num1, num2, ProjectClassTMP})
	}

	DarkTemplate := ChekDarkTemplate(r)

	lists := pro{template.HTML(bluemonday.UGCPolicy().SanitizeBytes([]byte("Project"))), ProjectTMP, DarkTemplate}

	templatesProject.Execute(w, lists)

}
