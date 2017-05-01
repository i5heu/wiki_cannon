package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Global sql.DB to access the database by all handlers
var db *sql.DB
var err error
var HtmlStructHeader string = `./template/header.html`
var HtmlStructFooter string = `./template/footer.html`

var cwd, _ = os.Getwd()
var fs = http.FileServer(http.Dir("static"))

func main() {
	// Create an sql.DB and check for errors
	db, err = sql.Open("mysql", "USER:PASSWORD@/blog_cannon")
	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS `article` (id INT NOT NULL AUTO_INCREMENT, `timec` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `needlogin` BOOLEAN NULL DEFAULT TRUE, `namespace` VARCHAR (128) NOT NULL DEFAULT 'main'  ,`title` VARCHAR (128) NOT NULL DEFAULT 'NO TITLE', `text` longtext,`tags` text, PRIMARY KEY (id),FULLTEXT(title,text,tags,namespace),FULLTEXT INDEX (title,text));")

	http.HandleFunc("/desk/", DesktopHandler)
	http.HandleFunc("/newentry", NewentryHandler)
	http.HandleFunc("/p/", ViewHandler)
	http.HandleFunc("/s/", SearchHandler)
	http.HandleFunc("/e/", EditHandler)
	http.HandleFunc("/api/", ApiHandler)

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", IndexHandler2)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler2(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/desk", 302)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
