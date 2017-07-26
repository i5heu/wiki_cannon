package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/microcosm-cc/bluemonday"
)

// Global sql.DB to access the database by all handlers
var db *sql.DB
var err error
var HtmlStructHeader string = `./template/header.html`
var HtmlStructFooter string = `./template/footer.html`

var cwd, _ = os.Getwd()
var fs = http.FileServer(http.Dir("static"))
var WcVersionUpdate string = ""
var WcVersionUpdateBOOL bool = false

func main() {
	// Create an sql.DB and check for errors
	db, err = sql.Open("mysql", dblogin)
	db.SetConnMaxLifetime(time.Second * 2)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic
	}

	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS `article` (`id` int(11) NOT NULL AUTO_INCREMENT,	 `timec` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,	 `timelastedit` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,	 `needlogin` tinyint(1) DEFAULT '1',	 `namespace` varchar(128) NOT NULL DEFAULT 'main',	 `title` varchar(128) NOT NULL DEFAULT 'NO TITLE',	 `text` longtext,	 `tags` text,	 `viewcounter` int(11) DEFAULT NULL,	 `editcounter` int(11) DEFAULT NULL,	 PRIMARY KEY (`id`),	 FULLTEXT KEY `title` (`title`,`text`,`tags`,`namespace`),	 FULLTEXT KEY `title_2` (`title`,`text`)	) ENGINE=InnoDB DEFAULT CHARSET=latin1")

	db.Exec("CREATE TABLE IF NOT EXISTS `items` ( `ItemID` int NOT NULL AUTO_INCREMENT, `timecreate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `timelastedit` timestamp NULL DEFAULT NULL,`needlogin` tinyint(1) DEFAULT '1',`APP` varchar(20) DEFAULT NULL,`viewcounter` INT NULL,`editcounter` int(11) DEFAULT NULL,`title1` varchar(128) DEFAULT NULL,`title2` varchar(128) NOT NULL,`text1` text,`text2` text,`tags1` varchar(256) DEFAULT NULL,`num1` int(64) DEFAULT NULL,`num2` int(64) DEFAULT NULL,`num3` int(64) DEFAULT NULL,PRIMARY KEY (`ItemID`)) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1")

	db.Exec("CREATE TABLE IF NOT EXISTS `eventlog` ( `id` int(11) NOT NULL AUTO_INCREMENT, `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `name` varchar(256) DEFAULT NULL, `changeAPP` varchar(128) DEFAULT NULL, `changeID` int(11) DEFAULT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=latin1")

	db.Exec("CREATE TABLE IF NOT EXISTS `BUarticle` ( `BUID` int(11) NOT NULL AUTO_INCREMENT, `BUtimestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `id` int(11) NOT NULL, `timec` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `timelastedit` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `needlogin` tinyint(1) DEFAULT '1', `namespace` varchar(128) NOT NULL DEFAULT 'main', `title` varchar(128) NOT NULL DEFAULT 'NO TITLE',`text` longtext, `tags` text, `viewcounter` int(11) DEFAULT NULL,`editcounter` int(11) DEFAULT NULL, PRIMARY KEY (`BUID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1")

	db.Exec("CREATE TABLE `BUitems` ( `BUItemID` int(11) NOT NULL AUTO_INCREMENT, `ItemID` int(11) DEFAULT NULL, `timecreate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `timelastedit` timestamp NULL DEFAULT NULL, `needlogin` tinyint(1) DEFAULT '1', `APP` varchar(20) DEFAULT NULL, `viewcounter` int(11) DEFAULT NULL, `editcounter` int(11) DEFAULT NULL, `title1` varchar(128) DEFAULT NULL, `title2` varchar(128) NOT NULL, `text1` text, `text2` text, `tags1` varchar(256) DEFAULT NULL, `num1` int(64) DEFAULT NULL, `num2` int(64) DEFAULT NULL, `num3` int(64) DEFAULT NULL, PRIMARY KEY (`BUItemID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1")

	fmt.Println("START")
	go func() {
		for {
			TMPCACHEWRITE = true
			time.Sleep(500 * time.Millisecond)
			Geldlogfunc("geldlog-false")
			Geldlogfunc("geldlog-true")
			cache(false, "article-false")
			cache(true, "article-true")
			Namespacefunc("namespace-true")
			Namespacefunc("namespace-false")
			Eventlogfunc("event-false")
			Eventlogfunc("event-true")
			Projectfunc("project-false")
			Projectfunc("project-true")
			Lasteditfunc("lastedit-false")
			Lasteditfunc("lastedit-true")
			TMPCACHEWRITE = false
			time.Sleep(500 * time.Millisecond)

			TMPCACHECACHEWRITE = true
			time.Sleep(500 * time.Millisecond)

			for key, value := range TMPCACHE {
				TMPCACHECACHE[key] = value
			}

			TMPCACHECACHEWRITE = false

			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			timeout := time.Duration(5 * time.Second)
			client := http.Client{
				Timeout: timeout,
			}
			resp, err := client.Get("https://i5heu.github.io/wiki_cannon/curentversion.html")
			checkErr(err)
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			checkErr(err)

			CurentVersionReponse := string(body)

			CurentVersionReponse = strings.Join(strings.Fields(CurentVersionReponse), " ")

			CurentVersionReponse = string(bluemonday.UGCPolicy().SanitizeBytes([]byte(CurentVersionReponse)))

			if CurentVersionReponse != wcversion {
				WcVersionUpdate = CurentVersionReponse
				WcVersionUpdateBOOL = true
			} else {
				WcVersionUpdateBOOL = false
			}
			time.Sleep(15 * time.Minute)
		}
	}()

	http.HandleFunc("/desk/", DesktopHandler)
	http.HandleFunc("/newentry", NewentryHandler)
	http.HandleFunc("/p/", ViewHandler)
	http.HandleFunc("/s/", SearchHandler)
	http.HandleFunc("/e/", EditHandler)
	http.HandleFunc("/project/", ProjectHandler)
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
		fmt.Println("\033[0;31m", err, "\033[0m")
	}
}
