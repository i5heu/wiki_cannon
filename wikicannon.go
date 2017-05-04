package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var guestmode bool = true

// Global sql.DB to access the database by all handlers
var db *sql.DB
var err error
var HtmlStructHeader string = `./template/header.html`
var HtmlStructFooter string = `./template/footer.html`

var cwd, _ = os.Getwd()
var fs = http.FileServer(http.Dir("static"))

func main() {
	// Create an sql.DB and check for errors
	db, err = sql.Open("mysql", "USER:PASSWORD@/wiki_cannon")
	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS `article` (id INT NOT NULL AUTO_INCREMENT, `timec` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, `needlogin` BOOLEAN NULL DEFAULT TRUE, `namespace` VARCHAR (128) NOT NULL DEFAULT 'main'  ,`title` VARCHAR (128) NOT NULL DEFAULT 'NO TITLE', `text` longtext,`tags` text, PRIMARY KEY (id),FULLTEXT(title,text,tags,namespace),FULLTEXT INDEX (title,text));")

	db.Exec("CREATE TABLE IF NOT EXISTS `items` ( `BUID` int(11) NOT NULL AUTO_INCREMENT, `id` int(11) DEFAULT NULL, `timecreate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `timelastedit` timestamp NULL DEFAULT NULL,`needlogin` tinyint(1) DEFAULT NULL,`APP` varchar(20) DEFAULT NULL,`editcounter` int(11) DEFAULT NULL,`title1` varchar(128) DEFAULT NULL,`title2` varchar(128) NOT NULL,`text1` text,`text2` text,`tags1` int(11) DEFAULT NULL,`num1` int(11) DEFAULT NULL,`num2` int(11) DEFAULT NULL,`num3` int(11) DEFAULT NULL,PRIMARY KEY (`BUID`)) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1")
	fmt.Println("START")

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
