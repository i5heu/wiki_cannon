package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type API2STRUCT struct {
	PWD       string
	APP       string
	title1    string
	title2    string
	text1     string
	text2     string
	tags1     string
	num1      int
	num2      int
	num3      int
	needlogin bool
	//ADDATIONAL
	string1 string
	string2 string
	string3 string
	int1    int
	int2    int
	int3    int
}

func ApiHandler2(w http.ResponseWriter, r *http.Request) { //THIS ONE IS WORKING WITH JSON Requests
	startAPI2 := time.Now()

	APILogin := false

	decoder := json.NewDecoder(r.Body)
	var json API2STRUCT
	errSearch := decoder.Decode(&json)
	if errSearch != nil {
		fmt.Fprintf(w, "ERROR")
		checkErr(err)
		return
	}

	if personalpwd == json.PWD {
		APILogin = true
	}

	switch json.APP {
	case "ItemWrite":
		ItemWrite(w, json)
	case "PwdManager":
		PwdManager(w, json)
	default:
		fmt.Fprintf(w, "ERROR")
	}

	fmt.Println("Api2Handler:", time.Since(startAPI2), APILogin)
}

func ItemWrite(w http.ResponseWriter, json API2STRUCT) {

	db.Exec("INSERT INTO items(APP,title1,title2,text1,text2,tags1,num1,num2,num3,needlogin) VALUES(?,?,?,?,?,?,?,?,?,?)", json.APP, json.title1, json.title2, json.text1, json.text2, json.tags1, json.num1, json.num2, json.num3, json.needlogin)
	fmt.Fprintf(w, `{"status":"OK"}`)
	return
}

func PwdManager(w http.ResponseWriter, json API2STRUCT) {

	var ids *sql.Rows

	ids, err = db.Query("SELECT title1,title2,text1 FROM items WHERE APP='PwdManager'")
	defer ids.Close()

	var jsonoutputtmp []string

	checkErr(err)
	for ids.Next() {
		var title1 string
		var title2 string
		var text1 string
		_ = ids.Scan(&title1, &title2, &text1)
		checkErr(err)

		jsonoutputtmp2 := `{ "title1" :"` + title1 + `" , "title2" :" ` + title2 + `","text1":"` + text1 + `" }`

		jsonoutputtmp = append(jsonoutputtmp, jsonoutputtmp2)
	}

	jsonoutput := `{ "PwdResult" : [` + strings.Join(jsonoutputtmp, ",") + "]}"

	fmt.Fprintf(w, jsonoutput)

	fmt.Println("PwdMa")
	return
}
