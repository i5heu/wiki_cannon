package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Global sql.DB to access the database by all handlers
var db *sql.DB
var err error
var HtmlStructHeader string = `header.html`
var HtmlStructFooter string = `footer.html`

func main() {
	// Create an sql.DB and check for errors
	db, err = sql.Open("mysql", "USER:PASSWORD@/blog_cannon")
	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS `article` (id INT NOT NULL AUTO_INCREMENT, `timec` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,`namespace` VARCHAR (128) NOT NULL DEFAULT 'main'  ,`title` VARCHAR (128) NOT NULL DEFAULT 'NO TITLE', `text` longtext,`tags` text, PRIMARY KEY (id),FULLTEXT(title,text,tags,namespace),FULLTEXT INDEX (title,text));")

	http.HandleFunc("/index/", IndexHandler)
	http.HandleFunc("/newentry", NewentryHandler)
	http.HandleFunc("/p/", ViewHandler)
	http.HandleFunc("/s/", SearchHandler)
	http.HandleFunc("/e/", EditHandler)
	http.HandleFunc("/api/", ApiHandler)
	http.HandleFunc("/", IndexHandler2)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler2(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", 302)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
