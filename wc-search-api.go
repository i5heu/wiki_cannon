package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

type SearchLoginSTRUCT struct {
	PWD         string
	SearchValue string
}

func ApiSearch(w http.ResponseWriter, r *http.Request) {
	startSearch := time.Now()

	SearchLogin := false

	decoder := json.NewDecoder(r.Body)
	var Sjson SearchLoginSTRUCT
	errSearch := decoder.Decode(&Sjson)
	if errSearch != nil {
		fmt.Fprintf(w, "ERROR")
		checkErr(err)
		return
	}

	if conf.AdminHASH == Sjson.PWD {
		SearchLogin = true
	}

	searchterm := ReplaceSpecialCharsWith_(Sjson.SearchValue)

	newquery := "%" + searchterm + "%"

	var ids *sql.Rows

	ids, err = db.Query("SELECT  id,namespace,title,tags,SUBSTR(text,1,150) FROM article WHERE (needlogin = '0' OR needlogin = ?) AND CONCAT(title,tags,namespace) LIKE ? ORDER BY timelastedit DESC LIMIT 200", SearchLogin, newquery)
	defer ids.Close()

	var jsonoutputtmp []string

	checkErr(err)
	for ids.Next() {
		var id int
		var namespace string
		var title string
		var tags string
		var text string
		_ = ids.Scan(&id, &namespace, &title, &tags, &text)
		checkErr(err)

		TitleTMP := string(bluemonday.UGCPolicy().SanitizeBytes([]byte(title)))
		TagsTMP := string(bluemonday.UGCPolicy().SanitizeBytes([]byte(tags)))
		TextTMP := string(bluemonday.UGCPolicy().SanitizeBytes([]byte(text)))

		jsonoutputtmp2 := `{ "id" : ` + strconv.Itoa(id) + ` , "namespace" :"` + namespace + `","title":"` + TitleTMP + `","tags":"` + TagsTMP + `","text":"` + ReplaceSpecialCharsWithSpaceSpaceALLOWED(TextTMP) + `" }`

		jsonoutputtmp = append(jsonoutputtmp, jsonoutputtmp2)
	}

	jsonoutput := `{ "SearchResult" : [` + strings.Join(jsonoutputtmp, ",") + "]}"

	fmt.Fprintf(w, jsonoutput)
	fmt.Println("SearchApi:", time.Since(startSearch), "Login:", SearchLogin)

}
