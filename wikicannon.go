package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/microcosm-cc/bluemonday"
)

//VERSION
var wcversion string = "WC-MR-V0.1"

// Global sql.DB to access the database by all handlers
var db *sql.DB
var err error
var HtmlStructHeader string
var HtmlStructFooter string

var cwd, _ = os.Getwd()
var fs = http.FileServer(http.Dir("static"))
var WcVersionUpdate string = ""
var WcVersionUpdateBOOL bool = false

var templatesDesktop, templatesEdit, templatesView, namespaceView, templatesProject, templatesSearch *template.Template

type Config struct {
	Dblogin                  string
	Guestmode                bool
	AdminPWD                 string
	GuestPWD                 string
	Templatefolder           string
	AdminHASH                string
	JabberNotification       bool
	JabberHost               string
	JabberUser               string
	JabberPassword           string
	JabberServerName         string
	JabberJIDreciever        string
	JabberInsecureSkipVerify bool
}

var conf Config

func main() {

	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
		return
	}

	var foo = sha256.Sum256([]byte(conf.AdminPWD)) //sha256 Parser for Password Token
	conf.AdminHASH = hex.EncodeToString(foo[:])

	HtmlStructHeader = conf.Templatefolder + `/header.html`
	HtmlStructFooter = conf.Templatefolder + `/footer.html`

	templatesDesktop = template.Must(template.ParseFiles("./template/desktop.html", HtmlStructHeader, HtmlStructFooter))
	templatesEdit = template.Must(template.ParseFiles("./template/edit.html", HtmlStructHeader, HtmlStructFooter))
	namespaceView = template.Must(template.ParseFiles("./template/namespace.html", HtmlStructHeader, HtmlStructFooter))
	templatesView = template.Must(template.ParseFiles("./template/view.html", HtmlStructHeader, HtmlStructFooter))
	templatesProject = template.Must(template.ParseFiles("./template/project.html", HtmlStructHeader, HtmlStructFooter))
	templatesSearch = template.Must(template.ParseFiles("./template/search.html", HtmlStructFooter, HtmlStructHeader))

	// ################ END CONFIG ###########################

	// Create an sql.DB and check for errors
	db, err = sql.Open("mysql", conf.Dblogin)
	db.SetConnMaxLifetime(time.Second * 2)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	checkErr(err)
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	CreateTable()

	fmt.Println("START")
	go func() {
		for {
			refreshCache()
			time.Sleep(10 * time.Minute)
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
			if err == nil {
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				checkErr(err)
				if err == nil {
					CurentVersionReponse := string(body)

					CurentVersionReponse = strings.Join(strings.Fields(CurentVersionReponse), " ")

					CurentVersionReponse = string(bluemonday.UGCPolicy().SanitizeBytes([]byte(CurentVersionReponse)))

					if CurentVersionReponse != wcversion {
						WcVersionUpdate = CurentVersionReponse
						WcVersionUpdateBOOL = true
					} else {
						WcVersionUpdateBOOL = false
					}
				}
			}
			time.Sleep(15 * time.Minute)
		}
	}()

	reciveXMPP()

	http.HandleFunc("/desk/", DesktopHandler)
	http.HandleFunc("/newentry", NewentryHandler)
	http.HandleFunc("/p/", ViewHandler)
	http.HandleFunc("/s/", SearchHandler)
	http.HandleFunc("/e/", EditHandler)
	http.HandleFunc("/project/", ProjectHandler)
	http.HandleFunc("/api/", ApiHandler)
	http.HandleFunc("/api2", ApiHandler2)
	http.HandleFunc("/api-search/", ApiSearch)

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/favicon.ico", FaviconHandler)
	http.HandleFunc("/", IndexHandler2)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler2(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/desk", 302)

}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/favicon/favicon.ico", 302)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println("\033[0;31m", err, "\033[0m")
		foo := "WikiERR:\n" + err.Error()
		sendXMPP(foo)
		err = nil
	}
}
